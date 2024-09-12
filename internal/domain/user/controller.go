package user

type userController struct {
	service UserServiceInterface
}

type UserControllerInterface interface {
}

func NewUserController(service UserServiceInterface) UserControllerInterface {
	return &userController{service}
}
