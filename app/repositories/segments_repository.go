package repositories

import (
	"fmt"
	"log"

    "gorm.io/gorm"
    "github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/models"
)

type SegmentsRepository struct {
	db *gorm.DB
}

func NewSegmentsRepository(db *gorm.DB) *SegmentsRepository {
	return &SegmentsRepository{
		db: db,
	}
}

func (sr *SegmentsRepository) Create(segment models.Segments) (models.Segments, error) {
	err := sr.db.Create(&segment).Error
	if err != nil {
		return models.Segments{}, err
	}
	return segment, err
}

func (sr *SegmentsRepository) FindByID(id uint) (models.Segments, error) {
	entity := models.Segments{
		ID: uint(id),
	}
	err := sr.db.First(&entity).Error
	if err != nil {
		return models.Segments{}, fmt.Errorf(`segment with id=%d not found: %w`, id, err)
	}
	return entity, err
}

func (sr *SegmentsRepository) FindByName(name string) (models.Segments, error) {
	entity := models.Segments{
		Name: name,
	}
	err := sr.db.First(&entity).Error
	if err != nil {
		return models.Segments{}, fmt.Errorf(`segment with name=%s not found: %w`, name, err)
	}
	return entity, err
}

func (sr *SegmentsRepository) RemoveSegment(name string) error {
	segment, err := sr.FindByName(name)
	if err != nil {
		return fmt.Errorf(`segment with name=%s not found: %w`, name, err)
	}

	var users []models.Users
	if err = sr.db.Joins("JOIN users_segments on users_segments.user_id=users.id").
			 Joins("JOIN segments on segments.id=users_segments.segment_id").
			 Where("segments.name=?", name).
			 Find(&users).Error; err != nil {
				log.Fatal(err)
				return fmt.Errorf(`error on finding segment with name=%s: %w`, name, err)
	 		 }
	for _, value := range users {
		fmt.Println("%w\n", value)
		err = models.RemoveSegment(&value, segment)
		if err != nil {
			return fmt.Errorf(`error when removing segment with name=%s on user with id=%d: %w`, name, value.ID, err)
		}
		sr.db.Save(&value)
	}
	sr.db.Delete(segment)
	return nil
}