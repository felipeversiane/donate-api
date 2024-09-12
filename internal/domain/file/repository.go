package file

import "github.com/felipeversiane/donate-api/internal/infra/database"

type fileRepository struct {
	db database.DatabaseInterface
}

type FileRepositoryInterface interface {
}

func NewFileRepository(db database.DatabaseInterface) FileRepositoryInterface {
	return &fileRepository{db}
}
