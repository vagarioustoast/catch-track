package models

import "gorm.io/gorm"

type Fisher struct {
	gorm.Model
	Name    string  `json:"name"`
	BoatID  uint    `gorm:"foreignKey:BoatID" json:"boatID"`
	Catches []Catch `gorm:"foreignKey:CatchID" json:"catches"`
}
