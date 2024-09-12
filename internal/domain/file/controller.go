package file

type fileController struct {
	service FileServiceInterface
}

type FileControllerInterface interface {
}

func NewFileController(service FileServiceInterface) FileControllerInterface {
	return &fileController{service}
}
