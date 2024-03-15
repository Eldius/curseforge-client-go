package persistence

import (
	"context"
	"fmt"
	"github.com/eldius/curseforge-client-go/internal/model"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type ModRepository struct {
	db *sqlx.DB
}

func NewModRepository(db *sqlx.DB) *ModRepository {
	return &ModRepository{db: db}
}

func (r *ModRepository) Save(ctx context.Context, m model.Mod) {
	_, err := r.db.NamedExecContext(ctx, insertModQuery, m.ToDBParams())
	if err != nil {
		err = fmt.Errorf("inserting mod data to local cache: %w", err)
		slog.With("error", err).Warn("ModRepository.Save")
	}
}
