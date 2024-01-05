package mocks

import (
	"context"
	"log-management/domain"

	"time"
)

// MockDAO is a struct that implements the DAO interface for testing
type MockDAO struct {
	Microservices *domain.Microservice
	Logs          *domain.Log
}

func (m *MockDAO) CreateMicroservice(ctx context.Context, microservice *domain.Microservice) (*domain.Microservice, error) {
	microservice.CreatedBy = "SYSTEM"
	microservice.UpdatedBy = "SYSTEM"

	return microservice, nil
}
func (m *MockDAO) CreateLog(ctx context.Context, log *domain.Log) (*domain.Log, error) {
	log.CreatedBy = "SYSTEM"
	log.UpdatedBy = "SYSTEM"

	return log, nil
}

func (m *MockDAO) GetMicroservice(ctx context.Context, id string) (*domain.Microservice, error) {
	microservice := &domain.Microservice{
		Base: domain.Base{
			ID:        id,
			CreatedBy: "SYSTEM",
			UpdatedBy: "SYSTEM",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:        "Test Microservice",
		Description: "Test Description",
	}

	return microservice, nil
}
func (m *MockDAO) GetLog(ctx context.Context, id string) (*domain.Log, error) {
	log := &domain.Log{
		Base: domain.Base{
			ID:        id,
			CreatedBy: "SYSTEM",
			UpdatedBy: "SYSTEM",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		MicroserviceID: "123",
		LogLevel:       "info",
		Message:        "Sample log message",
	}

	return log, nil
}

func (m *MockDAO) ListMicroservices(ctx context.Context) ([]*domain.Microservice, error) {
	microservices := []*domain.Microservice{
		{
			Base: domain.Base{
				ID:        "123",
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:        "TestMicroservice",
			Description: "Test Description",
		},
		{
			Base: domain.Base{
				ID:        "456",
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:        "TestMicroservice2",
			Description: "Test Description2",
		},
	}

	return microservices, nil
}
func (m *MockDAO) ListLogs(ctx context.Context, microserviceID string) ([]*domain.Log, error) {
	return []*domain.Log{m.Logs}, nil
}
func (m *MockDAO) ListAllLogs(ctx context.Context) ([]*domain.Log, error) {
	logs := []*domain.Log{
		{
			Base: domain.Base{
				ID:        "123",
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			MicroserviceID: "123",
			LogLevel:       "info",
			Message:        "Sample log message",
		},
		{
			Base: domain.Base{
				ID:        "456",
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			MicroserviceID: "123",
			LogLevel:       "info",
			Message:        "Sample log message",
		},
	}

	return logs, nil
}

func (m *MockDAO) UpdateMicroservice(ctx context.Context, microservice *domain.Microservice) (*domain.Microservice, error) {
	microservice.UpdatedAt = time.Now()
	microservice.UpdatedBy = "SYSTEM"

	return microservice, nil
}
func (m *MockDAO) UpdateLog(ctx context.Context, log *domain.Log) (*domain.Log, error) {
	log.UpdatedAt = time.Now()
	log.UpdatedBy = "SYSTEM"

	return log, nil
}

func (m *MockDAO) DeleteMicroservice(ctx context.Context, id string) error {
	return nil
}
func (m *MockDAO) DeleteLog(ctx context.Context, id string) error {
	return nil
}
