package models

type Product struct {
	Model
	Name             string    `json:"name,omitempty"`
	Price            int       `json:"price,omitempty"`
	Stock            int       `json:"stock,omitempty"`
	CustomAttributes string    `json:"customAttributes,omitempty"`
	Category         string    `json:"category,omitempty"`
	Barcodes         []Barcode `json:"barcodes,omitempty" gorm:"foreignkey:UUID"`
}