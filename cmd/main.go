package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fabiotavarespr/crm-backend/config"
	"github.com/fabiotavarespr/crm-backend/controllers"
	"github.com/fabiotavarespr/crm-backend/routes"
	"github.com/fabiotavarespr/crm-backend/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client

	customerCollection      *mongo.Collection
	customerService         services.CustomerService
	CustomerController      controllers.CustomerController
	CustomerRouteController routes.CustomerRouteController

	healthService         services.HealthService
	HealthController      controllers.HealthController
	HealthRouteController routes.HealthRouteController
)

func init() {

	// Load the .env variables
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Create a context
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Collections
	customerCollection = mongoclient.Database("crm").Collection("customers")
	customerService = services.NewCustomerService(customerCollection, ctx)
	CustomerController = controllers.NewCustomerController(customerService)
	CustomerRouteController = routes.NewCustomerRouteController(CustomerController)

	healthService = services.NewHealthService(mongoclient, ctx)
	HealthController = controllers.NewHealthController(healthService)
	HealthRouteController = routes.NewHealthRouteController(HealthController)

	// Create the Gin Engine instance
	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	CustomerRouteController.CustomerRoute(server)
	HealthRouteController.HealthRoute(server)

	log.Fatal(server.Run(":" + config.Port))
}
