package handler

import (
	"github.com/Koakovski/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitalizarHandler() {
	logger = config.GetLogger()
	db = config.GetDatabase()
}
