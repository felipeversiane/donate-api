package file

type fileService struct {
	repository FileRepositoryInterface
}

type FileServiceInterface interface {
}

func NewFileService(repository FileRepositoryInterface) FileRepositoryInterface {
	return &fileService{repository}
}
