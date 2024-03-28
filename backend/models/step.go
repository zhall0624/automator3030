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
	DB.Table(s.ActionType).Where("id = ?", s.ActionID).Take(&action)
	return action
}
