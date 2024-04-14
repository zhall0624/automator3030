package models

import (
	"time"

	"gorm.io/gorm"
)

type Step struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	ActionID      int
	ActionType    string
	SuccessStepID *int
	SuccessStep   *Step
	FailureStepID *int
	FailureStep   *Step
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (s *Step) Process() {
	s.Action().Run()
}

func (s *Step) Action() Action {
	var action Action

	switch s.ActionType {
	case "out_going_webhook_actions":
		action = new(OutGoingWebhookAction)
	default:
		panic("bad")
	}
	DB.Table(s.ActionType).Where("id = ?", s.ActionID).Take(&action)

	return action
}
