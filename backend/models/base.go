package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/satori/go.uuid"
)

type Model struct {
	ID        *uuid.UUID `sql:"type:uuid;primary_key;"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// BeforeCreate will set a UUID rather than numeric ID.
func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}
