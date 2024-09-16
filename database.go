package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	User        string
	Description string `json:"description"`
	Done        bool
}
type APITodo struct {
	Description string
	Done        bool
}

func database() *gorm.DB {
	var db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})
	return db
}
