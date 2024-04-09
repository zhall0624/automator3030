package models

import "gorm.io/gorm"

type Workflow struct {
	gorm.Model
	ID         int `param:"id" json:"id"`
	RootStepID int
	RootStep   Step
	Name       string `param:"name" json:"name"`
}

func (w *Workflow) ProcessSteps() {
	go w.RootStep.Process()
}
