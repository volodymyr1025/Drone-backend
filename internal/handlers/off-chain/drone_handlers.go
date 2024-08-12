package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"drone-backend/internal/database"
	"drone-backend/internal/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateDroneHandler - handler to create a new small drone
func CreateDroneHandler(w http.ResponseWriter, r *http.Request) {
	var drone models.Drone
	if err := json.NewDecoder(r.Body).Decode(&drone); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&drone).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(drone)
}

// GetDronesHandler - handler to get all small drones
func GetDronesHandler(w http.ResponseWriter, r *http.Request) {
	var drones []models.Drone
	if err := database.DB.Find(&drones).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drones)
}

// GetDroneHandler - handler to get a small drone by ID
func GetDroneHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var drone models.Drone
	if err := database.DB.First(&drone, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Drone not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drone)
}

// UpdateDroneHandler - handler to update a small drone by ID
func UpdateDroneHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var drone models.Drone
	if err := database.DB.First(&drone, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Drone not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&drone); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&drone).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drone)
}

// DeleteDroneHandler - handler to delete a small drone by ID
func DeleteDroneHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := database.DB.Delete(&models.Drone{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetDroneByZoneHandler(w http.ResponseWriter, r *http.Request) {
	zoneParam := mux.Vars(r)["zone"]
	zone, err := strconv.Atoi(zoneParam)

	if err != nil {
		http.Error(w, "Invalid zone parameter", http.StatusBadRequest)
		return
	}

	var drones []models.Drone
	if err := database.DB.Where("zone = ?", zone).Find(&drones).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drones)
}