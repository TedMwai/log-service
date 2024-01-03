package dao

import (
	"context"
	"log-management/domain"
)

type DAO interface {
	CreateMicroservice(ctx context.Context, microservice *domain.Microservice) (*domain.Microservice, error)
	CreateLog(ctx context.Context, log *domain.Log) (*domain.Log, error)

	GetMicroservice(ctx context.Context, id string) (*domain.Microservice, error)
	GetLog(ctx context.Context, id string) (*domain.Log, error)

	ListMicroservices(ctx context.Context) ([]*domain.Microservice, error)
	ListLogs(ctx context.Context, microserviceID string) ([]*domain.Log, error)

	UpdateMicroservice(ctx context.Context, microservice *domain.Microservice) (*domain.Microservice, error)
	UpdateLog(ctx context.Context, log *domain.Log) (*domain.Log, error)

	DeleteMicroservice(ctx context.Context, id string) error
	DeleteLog(ctx context.Context, id string) error
}