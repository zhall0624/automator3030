package models

import "gorm.io/gorm"

type Workflow struct {
	gorm.Model
	Name       string `json:"name"`
	RootStepID int
	RootStep   Step
}

func (w *Workflow) ProcessSteps() {
	go w.RootStep.Process()
}
