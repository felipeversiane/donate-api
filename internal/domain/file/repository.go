package file

import (
	"github.com/felipeversiane/donate-api/internal/infra/services/cloud"
	"github.com/felipeversiane/donate-api/internal/infra/services/database"
)

type fileRepository struct {
	db            database.DatabaseInterface
	objectStorage cloud.ObjectStorageInterface
}

type FileRepositoryInterface interface {
}

func NewFileRepository(
	db database.DatabaseInterface,
	objectStorage cloud.ObjectStorageInterface,
) FileRepositoryInterface {
	return &fileRepository{
		db:            db,
		objectStorage: objectStorage}
}
