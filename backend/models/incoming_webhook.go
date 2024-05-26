package models

import (
	"time"

	"gorm.io/gorm"
)

type IncomingWebHook struct {
	ID                      int    `gorm:"primaryKey" json:"id"`
	Name                    string `json:"name"`
	IncomingWebHookPayloads []IncomingWebHookPayload
	CreatedAt               time.Time      `json:"createdAt"`
	UpdatedAt               time.Time      `json:"updatedAt"`
	DeletedAt               gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
