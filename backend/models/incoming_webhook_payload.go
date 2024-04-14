package models

import (
	"time"

	"gorm.io/gorm"
)

type IncomingWebHookPayload struct {
	ID                int `gorm:"primaryKey" json:"id"`
	IncomingWebHookID int
	IncomingWebHook   IncomingWebHook
	Payload           string
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
