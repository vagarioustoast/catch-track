package models

import "gorm.io/gorm"

type Boat struct {
	gorm.Model
	Name     string  `json:"name"`
	FisherID uint    `gorm:"foreignKey:FisherID"`
	Catches  []Catch `gorm:"foreignKey:CatchID" json:"catches"`
}
