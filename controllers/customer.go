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

	ctx.JSON(http.StatusCreated, newCustomer)
}

func (cc *CustomerController) GetCustomers(ctx *gin.Context) {

	newCustomers, err := cc.customerService.GetCustomers()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if len(newCustomers.Customers) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "There's no customer"})
		return
	}

	ctx.JSON(http.StatusOK, newCustomers)
}

func (cc *CustomerController) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	newCustomer, err := cc.customerService.GetCustomer(id)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Customer not found"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newCustomer)
}

func (cc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	err := cc.customerService.DeleteCustomer(id)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Customer not found"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (cc *CustomerController) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	var customer *models.CustomerUpdateRequest

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updateCustomer, err := cc.customerService.UpdateCustomer(id, customer)

	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Customer not found"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, updateCustomer)
}
