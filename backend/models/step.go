package models

import (
	"gorm.io/gorm"
)

type Step struct {
	gorm.Model
	Name          string `json:"name"`
	ActionID      int
	ActionType    string
	SuccessStepID *int
	SuccessStep   *Step
	FailureStepID *int
	FailureStep   *Step
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
