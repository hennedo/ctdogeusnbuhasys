package models

import "github.com/gobuffalo/uuid"

type Log struct {
	Id      uuid.UUID `json:"id,omitempty" db:"id"`
	User    string    `json:"user,omitempty"`
	Product string    `json:"product,omitempty"`
	Action  string    `json:"action,omitempty"`
	Amount  int32     `json:"amount,omitempty"`
}
