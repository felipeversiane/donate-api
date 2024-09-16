package file

import (
	"github.com/felipeversiane/donate-api/internal/infra/services/aws"
	"github.com/felipeversiane/donate-api/internal/infra/services/database"
)

type fileRepository struct {
	db            database.DatabaseInterface
	objectStorage aws.ObjectStorageInterface
}

type FileRepositoryInterface interface {
}

func NewFileRepository(
	db database.DatabaseInterface,
	objectStorage aws.ObjectStorageInterface,
) FileRepositoryInterface {
	return &fileRepository{
		db:            db,
		objectStorage: objectStorage}
}
