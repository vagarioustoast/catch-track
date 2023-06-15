package models

import (
	"gorm.io/gorm"
)

type FishingMethodType string

const (
	// Fishing Methods
	Line FishingMethodType = "Line"
	Net  FishingMethodType = "Net"
)

type Catch struct {
	gorm.Model
	FisherID       uint              `gorm:"foreignKey:FisherID" json:"fisherID"`
	BoatID         uint              `gorm:"foreignKey:BoatID" json:"boatID"`
	SeafoodID      uint              `gorm:"foreignKey:SeafoodItemID" json:"seafoodID"`
	FishingLicense string            `json:"fishing_license"`
	CatchDateTime  string            `json:"catch_date_time"`
	CatchLocation  string            `json:"catch_location"`
	Quantity       int               `json:"quantity"`
	FishingMethod  FishingMethodType `json:"fishing_method"`
}
