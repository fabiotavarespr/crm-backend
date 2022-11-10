package controllers

import (
	"net/http"
	"strings"

	"github.com/fabiotavarespr/crm-backend/models"
	"github.com/fabiotavarespr/crm-backend/services"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return CustomerController{customerService}
}

func (cc *CustomerController) AddCustomer(ctx *gin.Context) {
	var customer *models.CustomerCreateRequest

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	newCustomer, err := cc.customerService.AddCustomer(customer)

	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"customer": newCustomer}})
}

func (cc *CustomerController) ListCustomers(ctx *gin.Context) {

	newCustomers, err := cc.customerService.ListCustomers()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": newCustomers})
}
