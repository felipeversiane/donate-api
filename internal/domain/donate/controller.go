package donate

type donateController struct {
	service DonateServiceInterface
}

type DonateControllerInterface interface {
}

func NewDonateController(service DonateServiceInterface) DonateControllerInterface {
	return &donateController{service}
}
