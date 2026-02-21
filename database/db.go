package database

import (
	"fmt"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(sqlite.Open("expense.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Database connected successfully")
}