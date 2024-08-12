package models

import (
	"gorm.io/gorm"
)

type Policy struct {
	gorm.Model
	Zone int	`json:"zone"`
	StartTime CustomTime	`json:"start_time"`
	EndTime CustomTime	`json:"end_time"`
}
