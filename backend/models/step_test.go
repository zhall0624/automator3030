package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStepActionAssociation(t *testing.T) {
	ConnectDatabase()
	step := Step{Name: "test"}
	action := OutGoingWebhookAction{Url: "test", Step: step}
	err := DB.Create(&action).Error

	if err != nil {
		panic("db error")
	}
	assert.Equal(t, step.Action(), action, "error message %", "formatted")
}
