package dao

import (
	"context"
	"database/sql"
	"errors"
	lg "github.com/rs/zerolog/log"
	"log-management/domain"
)

func (d *pgDAO) CreateLog(ctx context.Context, log *domain.Log) (*domain.Log, error) {
	if err := d.DB.NewInsert().Model(log).Returning("*").Scan(ctx); err != nil {
		// Scan return sql.ErrNoRows when no data is returned, unlike Exec
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		lg.Error().Err(err).Msg("failed to create log")
		return nil, err
	}

	return log, nil
}

// Get log by ID
func (d *pgDAO) GetLog(ctx context.Context, id string) (*domain.Log, error) {
	var log domain.Log
	if err := d.DB.NewSelect().Model(&log).Where("id = ?", id).Scan(ctx); err != nil {
		lg.Error().Err(err).Msg("failed to get log")
		return nil, err
	}

	return &log, nil
}

// List logs by microservice ID
func (d *pgDAO) ListLogs(ctx context.Context, microserviceID string) ([]*domain.Log, error) {
	var logs []*domain.Log
	if err := d.DB.NewSelect().Model(&logs).Where("microservice_id = ?", microserviceID).Scan(ctx); err != nil {
		lg.Error().Err(err).Msg("failed to list logs")
		return nil, err
	}

	return logs, nil
}

// Update log
func (d *pgDAO) UpdateLog(ctx context.Context, log *domain.Log) (*domain.Log, error) {
	if _, err := d.DB.NewUpdate().Model(log).Where("id = ?", log.ID).Exec(ctx); err != nil {
		lg.Error().Err(err).Msg("failed to update log")
		return nil, err
	}

	return log, nil
}

// Delete log
func (d *pgDAO) DeleteLog(ctx context.Context, id string) error {
	if _, err := d.DB.NewDelete().Model(&domain.Log{}).Where("id = ?", id).Exec(ctx); err != nil {
		lg.Error().Err(err).Msg("failed to delete log")
		return err
	}

	return nil
}
