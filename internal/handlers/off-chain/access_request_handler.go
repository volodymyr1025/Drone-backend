package handlers

import (
	"drone-backend/internal/database"
	"drone-backend/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var request models.AccessRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	requestSentence := fmt.Sprintf("Small Drone %s sent the access request", request.EntityID)

	accessJSON, err := json.Marshal(models.Access{
		Request: requestSentence,
		Status:  "Received",
	})
	if err != nil {
		log.Printf("Error marshaling Access: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Initialize history record with Access JSON
	history := models.History{
		Drone:  request.EntityID,
		Access: datatypes.JSON(accessJSON),
		Status: "In Progress",
	}

	if err := database.DB.Create(&history).Error; err != nil {
		log.Printf("Error creating history record: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// PIP
	drone, err := findDroneByID(request.EntityID)
	if err != nil {
		log.Printf("Error finding drone: %v", err)
		http.Error(w, "Drone does not exist.", http.StatusBadRequest)
	}
	attributes, pipContent := gatherAttributes(request.EntityID)
	pipJSON, err := json.Marshal(models.PIP{
		Content: pipContent,
		Status:  "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PIP: %v", err)
	}

	history.PIP = datatypes.JSON(pipJSON)

	// Update history record with PIP status
	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PIP: %v", err)
	}

	log.Printf("Gathered attributes: %+v", attributes)
	// PRP
	policies, prpContent := retrievePolicies(drone.Zone)
	prpJSON, err := json.Marshal(models.PRP{
		Content: prpContent,
		Status:  "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PRP: %v", err)
	}

	history.PRP = datatypes.JSON(prpJSON)

	// Update history record with PRP status
	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PRP: %v", err)
	}

	log.Printf("Retrieved policies: %+v", policies)

	// PDP
	accessGranted := evaluateAccess(attributes, policies);

	pdpDecision := "Denied"
	if accessGranted {
		pdpDecision = "Granted"
	}

	pdpJSON, err := json.Marshal(models.PDP{
		Decision: pdpDecision,
		Status:   "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PDP: %v", err)
	}

	history.PDP = datatypes.JSON(pdpJSON)
	history.Status = pdpDecision

	// Update history record with PDP status and overall decision
	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PDP: %v", err)
	}
	
	response := models.AccessResponse{
		Granted: accessGranted,
		Message: "Access request processed",
	}

	if !accessGranted {
		response.Message = "Access denied"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func findDroneByID(droneID string) (*models.Drone, error) {
	var drone models.Drone
	result := database.DB.First(&drone, "ID = ?", droneID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &drone, nil
}

func gatherAttributes(droneID string) (map[string]interface{}, string) {
	var drone models.Drone
	result := db.First(&drone, "id = ?", droneID)
	if result.Error != nil {
		return map[string]interface{}{
			"error": "Drone not found",
		}, "Drone not found"
	}
	
	loc, err := time.LoadLocation("Asia/Riyadh")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}

	attributes := map[string]interface{}{
		"drone_id": droneID,
		"time": time.Now().In(loc),
		"model_type": drone.ModelType,
		"zone": drone.Zone,
	}
	pipContent := fmt.Sprintf("Drone %s attributes gathered: Model Type: %s, Zone: %d, Time: %s",
		droneID, drone.ModelType, drone.Zone, attributes["time"].(time.Time).Format(time.RFC3339))

		return attributes, pipContent
}

func retrievePolicies(zone int) ([]models.Policy, string) {
	var policies []models.Policy
	result := db.Where("zone = ?", zone).Find(&policies)
	if result.Error != nil {
		return nil, "Failed to retrieve policies"
	}

	prpContent := fmt.Sprintf("Policies retrieved for zone %d policies found.",
		zone)
	if len(policies) > 0 {
		prpContent += " Policy details: "
		for i, policy := range policies {
			prpContent += fmt.Sprintf("[%d] Start: %s, End: %s; ",
				i+1, policy.StartTime, policy.EndTime)
		}
	}

	return policies, prpContent
}

func evaluateAccess(attributes map[string]interface{}, policies []models.Policy) bool {
	currentTime, ok := attributes["time"].(time.Time)
	if !ok {
		return false
	}

	for _, policy := range policies {
		if compareTimes(currentTime, policy.StartTime, policy.EndTime) {
				return true
		}
	}

	return false
}

func compareTimes(currentTime time.Time, startTime, endTime models.CustomTime) bool {
	// Extract the current time components
	currentHour, currentMinute, currentSecond := currentTime.Clock()

	// Create time objects for start and end times on the same day as currentTime
	startTimeToday := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, currentTime.Location())
	endTimeToday := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, currentTime.Location())

	// Compare the current time with start and end times
	if (currentHour > startTimeToday.Hour() || (currentHour == startTimeToday.Hour() && currentMinute > startTimeToday.Minute()) || (currentHour == startTimeToday.Hour() && currentMinute == startTimeToday.Minute() && currentSecond >= startTimeToday.Second())) &&
		(currentHour < endTimeToday.Hour() || (currentHour == endTimeToday.Hour() && currentMinute < endTimeToday.Minute()) || (currentHour == endTimeToday.Hour() && currentMinute == endTimeToday.Minute() && currentSecond <= endTimeToday.Second())) {
		return true
	}

	return false
}
