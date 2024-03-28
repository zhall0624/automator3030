package models

import (
	"bytes"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type OutGoingWebhookAction struct {
	gorm.Model
	Url        string
	HTTPMethod string
	Payload    string
	Step       Step `gorm:"polymorphic:Action"`
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
