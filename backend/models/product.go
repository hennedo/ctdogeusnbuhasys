package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Product struct {
	Id               uuid.UUID `json:"id,omitempty" db:"id"`
	CreatedAt        time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	Name             string    `json:"name,omitempty" db:"name"`
	Price            int       `json:"price,omitempty" db:"price"`
	Stock            int       `json:"stock,omitempty" db:"stock"`
	CustomAttributes string    `json:"customAttributes,omitempty" db:"customAttributes"`
	Category         string    `json:"category,omitempty"`
	Barcodes         []Barcode `json:"barcodes,omitempty" has_many:"barcodes"`
}