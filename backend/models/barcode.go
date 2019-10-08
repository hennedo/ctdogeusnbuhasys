package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Barcode struct {
	Id        uuid.UUID `json:"id,omitempty" db:"id"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	Ean       string    `json:"ean,omitempty" db:"ean"`
	ProductId uuid.UUID `json:"-" db:"product_id"`
	Product   *Product  `json:"product,omitempty" belongs_to:"product"`
}