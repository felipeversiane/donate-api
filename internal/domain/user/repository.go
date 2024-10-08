package user

import (
	"github.com/felipeversiane/donate-api/internal/infra/services/database"
)

type userRepository struct {
	db database.DatabaseInterface
}

type UserRepositoryInterface interface {
}

func NewUserRepository(db database.DatabaseInterface) UserRepositoryInterface {
	return &userRepository{db}
}
