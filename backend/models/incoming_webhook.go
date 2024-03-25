package models

import "gorm.io/gorm"

type IncomingWebHook struct {
	gorm.Model
	Name                   string `json:"name"`
	IncomingWebHookPayload []IncomingWebHookPayload
}
