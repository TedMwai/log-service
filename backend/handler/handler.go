package handler

import (
	"log-management/dao"

	"github.com/uptrace/bun"
)

type Handler struct {
	pg *bun.DB
	db dao.DAO
}

func New(d dao.DAO, pg *bun.DB) *Handler {
	return &Handler{
		pg: pg,
		db: d,
	}
}
