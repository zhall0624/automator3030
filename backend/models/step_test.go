package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestStepActionAssociation(t *testing.T) {
	ConnectDatabase()
	var step Step
	DB.Transaction(func(tx *gorm.DB) error {
		action := OutGoingWebhookAction{Url: "test", Step: Step{Name: "test"}}
		err := DB.Create(&action).Error

		DB.Model(&Step{}).Last(&step)
		if err != nil {
			panic("db error")
		}
		tx.Rollback()
		return nil
	})
	assert.NotNil(t, step.Action())
}
