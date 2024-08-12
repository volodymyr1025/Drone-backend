package models

import (
	"gorm.io/datatypes"
)

type Access struct {
	Request string `json:"request"`
	Status  string `json:"status"`
}

type PIP struct {
	Content string `json:"content"`
	Status  string `json:"status"`
}

type PRP struct {
	Content string `json:"content"`
	Status  string `json:"status"`
}

type PDP struct {
	Decision string `json:"decision"`
	Status   string `json:"status"`
}

type History struct {
	ID     uint           `json:"id" gorm:"primaryKey"`
	Drone  string         `json:"drone_id"`
	Access datatypes.JSON `json:"access"`
	PIP    datatypes.JSON `json:"pip"`
	PRP    datatypes.JSON `json:"prp"`
	PDP    datatypes.JSON `json:"pdp"`
	Status string         `json:"status"`
}
