package models

import "gorm.io/gorm"

type Retailer struct {
	gorm.Model
	Name         string    `json:"name"`
	SaleDateTime string    `json:"saleDateTime"`
	SaleLocation string    `json:"saleLocation"`
	SeafoodIDs   []Seafood `gorm:"foreignKey:SeafoodID" json:"seafoodIDs"`
}
