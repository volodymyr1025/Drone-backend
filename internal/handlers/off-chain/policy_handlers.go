package handlers

import (
	"encoding/json"
	"net/http"

	"drone-backend/internal/database"
	"drone-backend/internal/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreatePolicyHandler - handler to create a new policy
func CreatePolicyHandler(w http.ResponseWriter, r *http.Request) {
	var policy models.Policy
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&policy).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(policy)
}

// GetPoliciesHandler - handler to get all policies
func GetPoliciesHandler(w http.ResponseWriter, r *http.Request) {
	var policies []models.Policy
	if err := database.DB.Find(&policies).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(policies)
}

// GetPolicyHandler - handler to get a policy by ID
func GetPolicyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var policy models.Policy
	if err := database.DB.First(&policy, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Policy not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(policy)
}

// UpdatePolicyHandler - handler to update a policy by ID
func UpdatePolicyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var policy models.Policy
	if err := database.DB.First(&policy, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Policy not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&policy).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(policy)
}

// DeletePolicyHandler - handler to delete a policy by ID
func DeletePolicyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := database.DB.Delete(&models.Policy{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
