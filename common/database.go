package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"runGitDag/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "data/data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	PrepDb(db)

	return db
}

func PrepDb(db *gorm.DB) {
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.User{})
	}
	// Will add other checks for and create tables for other objects
}

func InitTestDb() *gorm.DB {
	test_db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	PrepDb(test_db)

	return test_db
}