package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger() *Logger {
	// Initialize Logger
	logger := NewLogger()
	return logger
}
