package donate

type donateService struct {
	repository DonateRepositoryInterface
}

type DonateServiceInterface interface {
}

func NewDonateService(repository DonateRepositoryInterface) DonateServiceInterface {
	return &donateService{repository}
}
