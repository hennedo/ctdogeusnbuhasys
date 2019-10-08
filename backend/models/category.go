package models

import (
	"github.com/gobuffalo/uuid"
	"time"
)

type Category struct {
	Id        uuid.UUID `json:"id,omitempty" db:"id"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	Name      string    `json:"name,omitempty" db:"name"`
}
