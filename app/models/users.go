package models

import (
	"fmt"

	gorm "gorm.io/gorm"
)

type UserEntity struct {
	ID uint `gorm:"primaryKey"`
	Segments []Segments `gorm:"many2many:users_segments"`
}