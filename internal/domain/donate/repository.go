package donate

import (
	"github.com/felipeversiane/donate-api/internal/infra/database"
)

type donateRepository struct {
	db database.DatabaseInterface
}

type DonateRepositoryInterface interface {
}

func NewDonateRepository(db database.DatabaseInterface) DonateRepositoryInterface {
	return &donateRepository{db}
}
