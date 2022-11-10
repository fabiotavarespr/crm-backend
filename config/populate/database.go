package populate

import (
	"github.com/fabiotavarespr/crm-backend/models"
	"github.com/fabiotavarespr/crm-backend/services"
	"github.com/jaswdr/faker"
)

type InitialDatabase struct {
	customerService services.CustomerService
}

func NewInitialDatabase(customerService services.CustomerService) InitialDatabase {
	return InitialDatabase{customerService}
}

func (id InitialDatabase) Start() error {
	fake := faker.New()
	for i := 0; i < 3; i++ {
		_, err := id.customerService.AddCustomer(&models.CustomerCreateRequest{
			Name:      fake.Person().Name(),
			Email:     fake.Person().Contact().Email,
			Role:      fake.Company().JobTitle(),
			Phone:     fake.Person().Contact().Phone,
			Contacted: fake.Bool(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
