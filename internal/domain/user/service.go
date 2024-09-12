package user

type userService struct {
	repository UserRepositoryInterface
}

type UserServiceInterface interface {
}

func NewUserService(repository UserRepositoryInterface) UserControllerInterface {
	return &userService{repository}
}
