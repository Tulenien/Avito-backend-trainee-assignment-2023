package repositories

import (
	"fmt"

    "gorm.io/gorm"
    "github.com/Tulenien/Avito-backend-trainee-assignment-2023/app/models"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (ur *UsersRepository) Create(user models.Users) (models.Users, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return models.Users{}, err
	}
	return user, err
}

func (ur *UsersRepository) FindByID(id uint) (models.Users, error) {
	entity := models.Users{
		ID: uint(id),
	}
	err := ur.db.First(&entity).Error
	if err != nil {
		return models.Users{}, fmt.Errorf(`user with id=%d not found: %w`, id, err)
	}
	return entity, err
}