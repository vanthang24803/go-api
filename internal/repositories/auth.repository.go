package repositories

import (
	"app/pkg/helpers"
)

type AuthRepository interface {
	GetAll() helpers.Response
}
