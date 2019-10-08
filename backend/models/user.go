package models

import (
	"github.com/gobuffalo/uuid"
	"time"
)

// Verifizierte E-Mail ist Pflicht für 'Minuskonten'
// Die Pin des Nutzers.
type User struct {
	Id        uuid.UUID `json:"id,omitempty" db:"id"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	Name      string    `json:"name,omitempty" db:"name"`
	ImageUrl  string    `json:"imageUrl,omitempty" db:"-"`
	Email     string    `json:"email,omitempty" db:"mail"`
	Verified  bool      `json:"verified,omitempty" db:"verified"`
	Gravatar  bool      `json:"gravatar,omitempty" db:"gravatar"`
	Balance   int32     `json:"balance,omitempty" db:"balance"`
	Pin       string    `json:"pin,omitempty" db:"pin"`
}
