package userentity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	Uuid      uuid.UUID      `gorm:"primaryKey, column:uuid" json:"uuid"`
	Email     string         `gorm:"column:email;index" json:"email"`
	Enabled   bool           `gorm:"column:enabled" json:"enabled"`
	FirstName string         `gorm:"column:firstName" json:"firstName"`
	LastName  string         `gorm:"column:lastName" json:"lastName"`
	Password  string         `gorm:"column:password" json:"password"`
	CreatedAt time.Time      `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt" json:"deletedAt"`
}

type CreatedEvent struct {
	ExchangeName string    `json:"-"`
	Uuid         uuid.UUID `json:"uuid"`
}
