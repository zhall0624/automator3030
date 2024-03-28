package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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
