package models

import "gorm.io/gorm"

type Workflow struct {
	gorm.Model
	RootStepID int
	RootStep   Step
	Name       string `param:"name" json:"name"`
}

func (w *Workflow) ProcessSteps() {
	go w.RootStep.Process()
}
