package dao

import (
	"context"
	"database/sql"
	"errors"
	"log-management/domain"
	"time"

	"github.com/rs/zerolog/log"
)

func (d *pgDAO) CreateMicroservice(ctx context.Context, microservice *domain.Microservice) (*domain.Microservice, error) {
	if err := d.DB.NewInsert().Model(microservice).Returning("*").Scan(ctx); err != nil {
		// Scan return sql.ErrNoRows when no data is returned, unlike Exec
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Error().Err(err).Msg("failed to create microservice")
		return nil, err
	}

	return microservice, nil
}

// Get microservice by ID
func (d *pgDAO) GetMicroservice(ctx context.Context, id string) (*domain.Microservice, error) {
	var microservice domain.Microservice
	if err := d.DB.NewSelect().Model(&microservice).Where("id = ?", id).Scan(ctx); err != nil {
		log.Error().Err(err).Msg("failed to get microservice")
		return nil, err
	}

	return &microservice, nil
}

// List microservices
func (d *pgDAO) ListMicroservices(ctx context.Context) ([]*domain.Microservice, error) {
	var microservices []*domain.Microservice
	if err := d.DB.NewSelect().Model(&microservices).Scan(ctx); err != nil {
		log.Error().Err(err).Msg("failed to list microservices")
		return nil, err
	}

	return microservices, nil
}

// Update microservice
func (d *pgDAO) UpdateMicroservice(ctx context.Context, microservice *domain.Microservice) (*domain.Microservice, error) {
	microservice.UpdatedAt = time.Now()

	if _, err := d.DB.NewUpdate().Model(microservice).WherePK().
		Column("name", "description", "updated_at").
		Returning("*").Exec(ctx); err != nil {
		log.Error().Err(err).Msg("failed to update microservice")
		return nil, err
	}

	return microservice, nil
}

// Delete microservice
func (d *pgDAO) DeleteMicroservice(ctx context.Context, id string) error {
	if _, err := d.DB.NewDelete().Model(&domain.Microservice{}).Where("id = ?", id).Exec(ctx); err != nil {
		log.Error().Err(err).Msg("failed to delete microservice")
		return err
	}

	return nil
}
