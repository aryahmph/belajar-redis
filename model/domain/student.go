package domain

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `gorm:"notNull"`
	Nim  string `gorm:"unique"`
}
