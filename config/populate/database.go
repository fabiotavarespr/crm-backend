package populate

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fabiotavarespr/crm-backend/models"
	"github.com/fabiotavarespr/crm-backend/services"
)

type InitialDatabase struct {
	customerService services.CustomerService
}

func NewInitialDatabase(customerService services.CustomerService) InitialDatabase {
	return InitialDatabase{customerService}
}

func (id InitialDatabase) Start() error {
	for i := 0; i < 3; i++ {
		_, err := id.customerService.AddCustomer(&models.CustomerRequest{
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Role:      gofakeit.JobDescriptor(),
			Phone:     gofakeit.Phone(),
			Contacted: gofakeit.Bool(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
