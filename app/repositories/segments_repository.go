package repositories

import (
	"fmt"

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

func (ur *SegmentsRepository) Create(segment models.Segments) (models.Segments, error) {
	err := ur.db.Create(&segment).Error
	if err != nil {
		return models.Segments{}, err
	}
	return segment, err
}

func (ur *SegmentsRepository) FindByID(id uint) (models.Segments, error) {
	entity := models.Segments{
		ID: uint(id),
	}
	err := ur.db.First(&entity).Error
	if err != nil {
		return models.Segments{}, fmt.Errorf(`segment with id=%d not found: %w`, id, err)
	}
	return entity, err
}

func (ur *SegmentsRepository) FindByName(name string) (models.Segments, error) {
	entity := models.Segments{
		Name: name,
	}
	err := ur.db.First(&entity).Error
	if err != nil {
		return models.Segments{}, fmt.Errorf(`segment with name=%s not found: %w`, name, err)
	}
	return entity, err
}