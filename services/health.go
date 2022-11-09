package services

import (
	"context"

	"github.com/fabiotavarespr/crm-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type HealthService interface {
	CheckHealth() (*models.HealthResponse, error)
}

type HealthServiceImpl struct {
	client *mongo.Client
	ctx    context.Context
}

func NewHealthService(client *mongo.Client, ctx context.Context) HealthService {
	return &HealthServiceImpl{client, ctx}
}

func (hs *HealthServiceImpl) CheckHealth() (*models.HealthResponse, error) {
	newHealth := new(models.HealthResponse)
	newHealth.Http = "OK"
	if err := hs.client.Ping(hs.ctx, readpref.Primary()); err != nil {
		newHealth.Mongo = "Fail"
	}
	newHealth.Mongo = "OK"

	return newHealth, nil
}
