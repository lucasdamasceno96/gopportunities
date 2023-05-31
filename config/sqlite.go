package config

import (
	"gopportunities/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite()(*gorm.DB, error){
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// cheack if db file exists

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
	}
	if err != nil {
		return nil, err
	}
	file, err := os.Create(dbPath)
		if err != nil {
		return nil, err
	}
	file.Close()

	// create database

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil{
		logger.Errorf("Sqlite opening error: %v ", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil{
		logger.Errorf("Sqlite migration error : %v ", err)
		return nil, err
	}

	return db, nil
}
