package models

import (
	"fmt"

	gorm "gorm.io/gorm"
)

type userEntity struct {
	Id uint `gorm:"primaryKey"`
	Segment pgtype.JSONB `gorm:"type:jsonb;default:'[]'"`
}