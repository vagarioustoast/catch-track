package models

import "gorm.io/gorm"

type ProcessingMethodType string

const (
	Cleaned  ProcessingMethodType = "Cleaned"
	Filleted ProcessingMethodType = "Filleted"
	Packaged ProcessingMethodType = "Packaged"
	Other    ProcessingMethodType = "Other"
)

type ProcessingPlant struct {
	gorm.Model
	Name               string               `json:"name"`
	ProcessingDateTime string               `json:"processingDateTime"`
	ProcessingMethod   ProcessingMethodType `json:"processingMethod"`
	Packaging          string               `json:"packaging"`
	SeafoodIDs         []Seafood            `gorm:"foreignKey:SeafoodID" json:"seafoodIDs"`
}
