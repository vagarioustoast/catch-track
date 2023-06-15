package models

import "gorm.io/gorm"

type TransportMethodType string

const (
	Truck TransportMethodType = "Truck"
	Plane TransportMethodType = "Plane"
)

type Distributor struct {
	gorm.Model
	Name                 string              `json:"name"`
	DistributionDateTime string              `json:"distributionDateTime"`
	TransportMethod      TransportMethodType `json:"transportMethod"`
	DeliveryLocation     string              `json:"deliveryLocation"`
	SeafoodIDs           []Seafood           `gorm:"foreignKey:SeafoodID" json:"seafoodIDs"`
}
