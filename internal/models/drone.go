package models

import (
	"gorm.io/gorm"
)

type Drone struct {
	gorm.Model
	ModelType string	`json:"model_type"`
	Zone int	`json:"zone"`
}