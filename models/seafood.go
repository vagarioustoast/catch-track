// ItemID: Unique identifier for each seafood item/batch
// Species: Type of seafood
// CatchID: Related catch information
// Size: Size of the seafood item
// Weight: Weight of the seafood item
// Condition: Status of the seafood (e.g., alive, fresh, frozen, etc.)
package models

import "gorm.io/gorm"

type SpeciesType string

const (
	Shrimp  SpeciesType = "Shrimp"
	Crab    SpeciesType = "Crab"
	Lobster SpeciesType = "Lobster"
	Fish    SpeciesType = "Fish"
)

type Seafood struct {
	gorm.Model
	Species  SpeciesType `json:"species"`
	CatchID  uint        `gorm:"foreignKey:CatchID" json:"catchID"`
	Size     string      `json:"size"`
	Weight   string      `json:"weight"`
	Quantity uint        `json:"quantity"`
}
