package domain

import (
	"time"
)

type Base struct {
	ID        string            `bun:",pk"`
	CreatedBy string            `bun:",notnull,nullzero,default:'SYSTEM'"`
	UpdatedBy string            `bun:",notnull,nullzero,default:'SYSTEM'"`
	CreatedAt time.Time         `bun:"default:now(),notnull,nullzero"`
	UpdatedAt time.Time         `bun:"default:now(),notnull,nullzero"`
	DeletedAt time.Time         `bun:",soft_delete,nullzero"`
	Metadata  map[string]string `bun:",notnull,default:'{}'"`
}
