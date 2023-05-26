package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeDatabase()
	if err != nil {
		return fmt.Errorf("error initialize database: %v", err)
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	// Initialize Logger
	logger = newLogger()
	return logger
}
