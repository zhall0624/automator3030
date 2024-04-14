package models

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type OutGoingWebhookAction struct {
	ID         int `gorm:"primaryKey" json:"id"`
	Url        string
	HTTPMethod string
	Payload    string
	Step       Step           `gorm:"polymorphic:Action"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (a *OutGoingWebhookAction) Run() error {
	switch a.HTTPMethod {
	case "GET":
		_, err := http.Get(a.Url)
		return err
	case "POST":
		_, err := http.Post(a.Url, "json", bytes.NewBufferString(a.Payload))
		return err
	default:
		return errors.New("unexpected HTTP method for OutGoingWebhookAction, expected GET or POST")
	}
}
