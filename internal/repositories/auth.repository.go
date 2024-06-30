package repositories

import "app/internal/models"

type AuthRepository interface {
	GetAll() ([]models.User, error)
}
