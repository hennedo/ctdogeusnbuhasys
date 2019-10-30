package models

import (
	"github.com/gofrs/uuid"
)

type Barcode struct {
	Model
	Ean       string    `json:"ean,omitempty"`
	ProductId uuid.UUID `json:"-"`
	Product   *Product  `json:"product,omitempty"`
}
