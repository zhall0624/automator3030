package models

import (
	"time"

	"gorm.io/gorm"
)

type Workflow struct {
	ID         int            `gorm:"primaryKey" param:"id" json:"id"`
	RootStepID *int           `json:"rootStepId"`
	RootStep   Step           `json:"-"`
	Name       string         `param:"name" json:"name"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (w *Workflow) ProcessSteps() {
	go w.RootStep.Process()
}
