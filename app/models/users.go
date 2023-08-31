package models

import (
	"fmt"
)

type Users struct {
	ID uint `gorm:"primaryKey"`
	Segments []Segments `gorm:"many2many:users_segments"`
}

func AddSegment(user *Users, segment Segments) error {
	for _, value := range user.Segments {
		if value.Name == segment.Name {
			return fmt.Errorf(`User already has segment`)
		}
	}
	user.Segments = append(user.Segments, segment)
	return nil
}

func RemoveIndex(s []Segments, index int) ([]Segments, error) {
	if index < 0 || index >= len(s) {
		return nil, fmt.Errorf(`Index out of bounds`)
	} else {
		j := 0
		for i := range s {
			if i != index {
				s[j] = s[index]
				j++
			}
		}
		new_s := s[:j]
		return new_s, nil
	}
}

func RemoveSegment(user *Users, segment Segments) error {
	for index, value := range user.Segments {
		if value.Name == segment.Name {
			new_s, err := RemoveIndex(user.Segments, index)
			if err != nil {
				return err
			}
			user.Segments = new_s
			return nil
		}
	}
	return fmt.Errorf(`Not found`)
}