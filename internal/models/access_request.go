package models

type AccessRequest struct {
	EntityID string `json:"entity_id"`
}

type AccessResponse struct {
	Granted bool   `json:"granted"`
	Message string `json:"message"`
}