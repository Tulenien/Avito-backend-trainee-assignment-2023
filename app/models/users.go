package models

type Users struct {
	ID uint `gorm:"primaryKey"`
	Segments []Segments `gorm:"many2many:users_segments"`
}