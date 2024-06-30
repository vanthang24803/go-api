package impl

import (
	"app/internal/models"
	"app/internal/repositories"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	Db *gorm.DB
}

func (a *AuthRepositoryImpl) GetAll() ([]models.User, error) {
	var users []models.User

	if err := a.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func NewAuthRepositoryImpl(Db *gorm.DB) repositories.AuthRepository {
	return &AuthRepositoryImpl{Db: Db}
}
