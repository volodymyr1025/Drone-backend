package api

import (
	handler "drone-backend/internal/handlers/on-chain"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PolicyAPI struct {
	Handler *handler.PolicyHandler
}

func NewPolicyAPI(h *handler.PolicyHandler) *PolicyAPI {
	return &PolicyAPI{Handler: h}
}

type PolicyResponse struct {
	ID uint
	Zone int	`json:"zone"`
	StartTime string	`json:"start_time"`
	EndTime string	`json:"end_time"`
}

type CreatePolicyRequest struct {
	Zone int `json:"zone"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
}

type CreatePolicyResponse struct {
	ID	uint
	Zone int	`json:"zone"`
	StartTime string	`json:"start_time"`
	EndTime	string	`json:"end_time"`
}

type UpdatePolicyRequest struct {
	Zone int `json:"zone"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
}

type UpdatePolicyResponse struct {
	ID	uint
	Zone int	`json:"zone"`
	StartTime string	`json:"start_time"`
	EndTime	string	`json:"end_time"`
}

type RemovePolicyResponse struct {
	TransactionHash string `json:"transaction_hash"`
}

type GetPolicyByZoneRequest struct {
	Zone int `json:"zone"`
}

type GetPolicyByZoneResponse struct {
	ID	uint
	Zone int	`json:"zone"`
	StartTime string	`json:"start_time"`
	EndTime	string	`json:"end_time"`
}

func (api *PolicyAPI) GetPoliciesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	policies, err := api.Handler.GetPolicies()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get policies: %v", err), http.StatusInternalServerError)
		return
	}

	var response []PolicyResponse
	for _, policy := range policies {
		response = append(response, PolicyResponse{
			ID:	uint(policy.Id.Uint64()),
			Zone: int(policy.Zone.Int64()),
			StartTime: policy.StartTime,
			EndTime: policy.EndTime,
		})
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}

func (api *PolicyAPI) CreatePolicyHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req CreatePolicyRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    policy, err := api.Handler.CreatePolicy(req.Zone, req.StartTime, req.EndTime)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to create policy: %v", err), http.StatusInternalServerError)
        return
    }

    response := CreatePolicyResponse{
        ID: uint(policy.Id.Uint64()),
        Zone: int(policy.Zone.Int64()),
        StartTime: policy.StartTime,
        EndTime: policy.EndTime,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *PolicyAPI) UpdatePolicyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
    uint64ID, _ := strconv.ParseUint(id, 10, 0)
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req UpdatePolicyRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    policy, err := api.Handler.UpdatePolicy(uint(uint64ID), req.Zone, req.StartTime, req.EndTime)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to update policy: %v", err), http.StatusInternalServerError)
        return
    }

    response := UpdatePolicyResponse{
        ID: uint(policy.Id.Uint64()),
        Zone: int(policy.Zone.Int64()),
        StartTime: policy.StartTime,
        EndTime: policy.EndTime,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *PolicyAPI) RemovePolicyHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
    uint64ID, _ := strconv.ParseUint(id, 10, 0)
    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    txHash, err := api.Handler.RemovePolicy(uint(uint64ID))
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to remove policy: %v", err), http.StatusInternalServerError)
        return
    }

    response := RemovePolicyResponse{
        TransactionHash: txHash,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}

func (api *PolicyAPI) GetPolicyByZoneHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    zoneStr := r.URL.Query().Get("zone")
    if zoneStr == "" {
        http.Error(w, "Missing 'zone' query parameter", http.StatusBadRequest)
        return
    }

    zone, err := strconv.Atoi(zoneStr)
    if err != nil {
        http.Error(w, "Invalid 'zone' query parameter", http.StatusBadRequest)
        return
    }

    policy, err := api.Handler.GetPolicyByZone(zone)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get policy by zone: %v", err), http.StatusInternalServerError)
        return
    }

    response := GetPolicyByZoneResponse{
        ID:        uint(policy.Id.Uint64()),
        Zone:      int(policy.Zone.Int64()),
        StartTime: policy.StartTime,
        EndTime:   policy.EndTime,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
    }
}
