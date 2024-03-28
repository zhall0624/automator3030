package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	tables := [4]interface{}{&IncomingWebHook{}, &OutGoingWebhookAction{}, &Workflow{}, &Step{}}
	for _, table := range tables {
		err = database.AutoMigrate(table)
		if err != nil {
			return
		}
	}

	DB = database
}
