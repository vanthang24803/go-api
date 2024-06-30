package impl

import (
	"app/internal/models"
	"app/internal/repositories"
	"app/pkg/helpers"
	"net/http"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	Db *gorm.DB
}

func (a *AuthRepositoryImpl) GetAll() helpers.Response {
	var users []models.User

	if err := a.Db.Find(&users).Error; err != nil {
		return helpers.Send(http.StatusBadRequest, "Something went wrong!")
	}

	return helpers.Send(http.StatusOK, users)
}

func NewAuthRepositoryImpl(Db *gorm.DB) repositories.AuthRepository {
	return &AuthRepositoryImpl{Db: Db}
}
