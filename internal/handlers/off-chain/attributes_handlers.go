package handlers

import (
	"encoding/json"
	"net/http"

	"drone-backend/internal/database"
	"drone-backend/internal/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateAttributeHandler(w http.ResponseWriter, r *http.Request) {
	var attribute models.Attribute
	if err := json.NewDecoder(r.Body).Decode(&attribute); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Create(&attribute).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attribute)
}

func GetAttributesHandler(w http.ResponseWriter, r *http.Request) {
	var attributes []models.Attribute
	if err := database.DB.Find(&attributes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attributes)
}

func GetAttributeHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var attribute models.Attribute
	if err := database.DB.First(&attribute, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Attribute not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attribute)
}

func UpdateAttributeHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var attribute models.Attribute
	if err := database.DB.First(&attribute, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Attribute not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&attribute); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Save(&attribute).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attribute)
}

func DeleteAttributeHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := database.DB.Delete(&models.Attribute{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetAttributeByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	var attributes []models.Attribute
	if err := database.DB.Where("name = ?", name).Find(&attributes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(attributes)
}