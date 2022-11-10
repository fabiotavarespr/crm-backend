package routes

import (
	"github.com/fabiotavarespr/crm-backend/controllers"
	"github.com/gin-gonic/gin"
)

type CustomerRouteController struct {
	customerController controllers.CustomerController
}

func NewCustomerRouteController(customerController controllers.CustomerController) CustomerRouteController {
	return CustomerRouteController{customerController}
}

func (rc *CustomerRouteController) CustomerRoute(rg *gin.Engine) {
	rg.GET("/customers/:id", rc.customerController.GetCustomer)
	rg.GET("/customers", rc.customerController.GetCustomers)
	rg.POST("/customers", rc.customerController.AddCustomer)
}
