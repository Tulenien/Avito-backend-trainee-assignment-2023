package models

import (
	"fmt"

	gorm "gorm.io/gorm"
)

type Segment struct {
	Id uint `gorm:"primaryKey"`
	Name string
}