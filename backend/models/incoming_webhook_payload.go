package models

import "gorm.io/gorm"

type IncomingWebHookPayload struct {
	gorm.Model
	IncomingWebHookID int
	IncomingWebHook   IncomingWebHook
	Payload           string
}
