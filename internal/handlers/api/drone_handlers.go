package api

import (
	handler "drone-backend/internal/handlers/on-chain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DroneAPI struct {
	Handler *handler.DroneHandler
}

func NewDroneAPI(h *handler.DroneHandler) *DroneAPI {
	return &DroneAPI{Handler: h}
}

type DroneResponse struct {
	ID uint
	ModelType string	`json:"model_type"`
	Zone int `json:"zone"`
}

type CreateDroneRequest struct {
	ModelType string `json:"model_type"`
	Zone int `json:"zone"`
}

type CreateDroneResponse struct {
    ID    uint
    ModelType string `json:"model_type"`
    Zone  int    `json:"zone"`
}

type UpdateDroneRequest struct {
    ModelType string `json:"model_type"`
    Zone  int    `json:"zone"`
}

type UpdateDroneResponse struct {
    TransactionHash string `json:"transaction_hash"`
}

type RemoveDroneResponse struct {
    TransactionHash string `json:"transaction_hash"`
}

type GetDroneResponse struct {
    ID    uint
    ModelType string `json:"model_type"`
    Zone  int    `json:"zone"`
}

func (api *DroneAPI) GetDronesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    drones, err := api.Handler.GetDrones()
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drones: %v", err), http.StatusInternalServerError)
        return
    }

    var response []DroneResponse
    for _, drone := range drones {
        response = append(response, DroneResponse{
            ID:    uint(drone.Id.Uint64()),
            ModelType: drone.ModelType,
            Zone:  int(drone.Zone.Int64()),
        })
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}

func (api *DroneAPI) GetDronesByZoneHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	zone := mux.Vars(r)["zone"]
    intZone, _ := strconv.Atoi(zone)

    drones, err := api.Handler.GetDronesByZone(intZone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drones by name: %v", err), http.StatusInternalServerError)
        return
    }

	var response []DroneResponse
	for _, drone := range drones {
		response = append(response, DroneResponse{
			ID: uint(drone.Id.Uint64()),
			ModelType: drone.ModelType,
			Zone: int(drone.Zone.Int64()),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func (api *DroneAPI) CreateDroneHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req CreateDroneRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    drone, err := api.Handler.CreateDrone(req.ModelType, req.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to create drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := GetDroneResponse{
        ID:    uint(drone.Id.Uint64()),
        ModelType: drone.ModelType,
        Zone:  int(drone.Zone.Int64()),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *DroneAPI) UpdateDroneHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
    uint64ID, _ := strconv.ParseUint(id, 10, 0)
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req UpdateDroneRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    drone, err := api.Handler.UpdateDrone(uint(uint64ID), req.ModelType, req.Zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to update drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := GetDroneResponse{
        ID:    uint(drone.Id.Uint64()),
        ModelType: drone.ModelType,
        Zone:  int(drone.Zone.Int64()),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *DroneAPI) RemoveDroneHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
    uint64ID, _ := strconv.ParseUint(id, 10, 0)
    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    txHash, err := api.Handler.RemoveDrone(uint(uint64ID))
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to remove drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := RemoveDroneResponse{
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *DroneAPI) GetDroneHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
    uint64ID, _ := strconv.ParseUint(id, 10, 0)
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    drone, err := api.Handler.GetDrone(uint(uint(uint64ID)))
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get drone: %v", err), http.StatusInternalServerError)
        return
    }

    response := GetDroneResponse{
        ID:    uint(drone.Id.Uint64()),
        ModelType: drone.ModelType,
        Zone:  int(drone.Zone.Int64()),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}