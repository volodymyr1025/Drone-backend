package api

import (
	"drone-backend/internal/database"
	handler "drone-backend/internal/handlers/on-chain"
	"drone-backend/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AccessAPI struct {
	PDPHandler *handler.PDPHandler
    DroneAPI    *DroneAPI
    PolicyAPI   *PolicyAPI
}

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

// var mu sync.Mutex
var mu sync.RWMutex

func NewAccessAPI(pdpHandler *handler.PDPHandler, droneAPI *DroneAPI, policyAPI *PolicyAPI) *AccessAPI {
	return &AccessAPI{PDPHandler: pdpHandler, DroneAPI: droneAPI, PolicyAPI: policyAPI}
}

type AccessRequest struct {
	EntityID string `json:"entity_id"`
    RequestTarget uint `json:"request_target"`
}

type AccessResponse struct {
    Granted          bool   `json:"granted"`
    TransactionHash  string `json:"transaction_hash"`
}

func (api *AccessAPI) Layer1AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    // mu.Lock()
    // defer mu.Unlock()
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    
    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    uint64ID, err := strconv.ParseUint(req.EntityID, 10, 0)
    if err != nil {
        http.Error(w, "Invalid entity ID", http.StatusBadRequest)
        return
    }
    
	mu.Lock()
    requestSentence := fmt.Sprintf("Small Drone %s sent the access request", req.EntityID)
	accessJSON, err := json.Marshal(models.Access{
        Request: requestSentence,
		Status:  "Received",
	})
	if err != nil {
		mu.Unlock()
        log.Printf("Error marshaling Access: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
    
    history := models.History{
        Drone:  req.EntityID,
		Access: datatypes.JSON(accessJSON),
		Status: "In Progress",
	}

	if err := database.DB.Create(&history).Error; err != nil {
		mu.Unlock()
        log.Printf("Error creating history record: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	mu.Unlock()

    // PIP
	mu.RLock() 
    drone, err := findDroneByID(req.EntityID)
	mu.RUnlock()

    if err != nil {
		log.Printf("Error finding drone: %v", err)
		http.Error(w, "Drone does not exist.", http.StatusBadRequest)
	}

    loc, err := time.LoadLocation("Asia/Riyadh")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}

    currentTime := time.Now().In(loc);
	pipContent := fmt.Sprintf("Drone %d attributes gathered: Model Type: %s, Zone: %d, Time: %s",
		drone.ID, drone.ModelType, drone.Zone, currentTime.Format(time.RFC3339))

	mu.Lock()
	pipJSON, err := json.Marshal(models.PIP{
		Content: pipContent,
		Status:  "Completed",
	})
	if err != nil {
		mu.Unlock()
		log.Printf("Error marshaling PIP: %v", err)
	}

	history.PIP = datatypes.JSON(pipJSON)

	if err := database.DB.Save(&history).Error; err != nil {
		mu.Unlock()
		log.Printf("Error updating history with PIP: %v", err)
	}
	mu.Unlock()

    // PRP
	mu.RLock()
    var policy models.Policy
    result := database.DB.First(&policy, "zone = ?", drone.Zone)
	mu.RUnlock()

    if result.Error != nil {
        http.Error(w, "Failed to retrieve policies", http.StatusInternalServerError)
        return
    }

	mu.Lock()
    prpContent := fmt.Sprintf("Policies retrieved for zone %d policies found.", policy.Zone)

    prpJSON, err := json.Marshal(models.PRP{
		Content: prpContent,
		Status:  "Completed",
	})
	if err != nil {
		log.Printf("Error marshaling PRP: %v", err)
	}

	history.PRP = datatypes.JSON(prpJSON)
	if err := database.DB.Save(&history).Error; err != nil {
		log.Printf("Error updating history with PRP: %v", err)
	}
	mu.Unlock()

    	// PDP

    accessGranted := false;
    if compareTimes(currentTime, policy.StartTime, policy.EndTime) {
        accessGranted = true;
    }

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


    // PDP
    granted, txHash, err := api.PDPHandler.Layer1EvaluateAccess(uint(uint64ID), drone.ModelType, policy.Zone, policy.StartTime.String(), policy.EndTime.String(), accessGranted)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }

    response := AccessResponse{
        Granted:         granted,
        TransactionHash: txHash,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *AccessAPI) Layer2AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }    
    uint64ID, err := strconv.ParseUint(req.EntityID, 10, 0)
    if err != nil {
        http.Error(w, "Invalid entity ID", http.StatusBadRequest)
        return
    }
    

    // PIP
	mu.RLock() 
    drone, err := findDroneByID(req.EntityID)
	mu.RUnlock()

    if err != nil {
		log.Printf("Error finding drone: %v", err)
		http.Error(w, "Drone does not exist.", http.StatusBadRequest)
	}

    // PRP
	mu.RLock()
    var policy models.Policy
    result := database.DB.First(&policy, "zone = ?", drone.Zone)
	mu.RUnlock()

    if result.Error != nil {
        http.Error(w, "Failed to retrieve policies", http.StatusInternalServerError)
        return
    }

    // PDP
    granted, txHash, err := api.PDPHandler.Layer2EvaluateAccess(uint(uint64ID), drone.ModelType, policy.Zone, policy.StartTime.String(), policy.EndTime.String())
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }

    response := AccessResponse{
        Granted:         granted,
        TransactionHash: txHash,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *AccessAPI) Layer3AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    uint64ID, _ := strconv.ParseUint(req.EntityID, 10, 0)

    // PIP - Read operation
    mu.RLock()
    drone, err := findDroneByID(req.EntityID)
    mu.RUnlock()

    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drone: %v", err), http.StatusInternalServerError)
        return
    }

    // PDP - Write operation (interacting with the smart contract)
    granted, txHash, err := api.PDPHandler.Layer3EvaluateAccess(uint(uint64ID), drone.ModelType, drone.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }
    
    response := AccessResponse{
        Granted:         granted,
        TransactionHash: txHash,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *AccessAPI) Layer4AccessRequestHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req AccessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Lock for writing the initial history record
    uint64ID, _ := strconv.ParseUint(req.EntityID, 10, 0)

    // PDP - Write operation (interacting with the smart contract)
    granted, txHash, err := api.PDPHandler.Layer4EvaluateAccess(uint(uint64ID))
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to evaluate access: %v", err), http.StatusInternalServerError)
        return
    }

    response := AccessResponse{
        Granted:         granted,
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func findDroneByID(droneID string) (*models.Drone, error) {
	var drone models.Drone
	result := database.DB.First(&drone, "ID = ?", droneID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &drone, nil
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