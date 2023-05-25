package schema

import (
	"gorm.io/gorm"
)

type OpeningSchema struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   uint64
}
