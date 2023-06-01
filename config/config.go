package config

import (
	"fmt"

	"gorm.io/gorm"
)

var(
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error 

	// initialize sqlite 

	db, err = InitializeSQlite()
	if err != nil {
		return fmt.Errorf("Error with sqlite initialize:  %v", err)
	}


	return nil
}

func GetSQlite() *gorm.DB {
	return db
 }

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
