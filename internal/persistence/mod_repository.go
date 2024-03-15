package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eldius/curseforge-client-go/internal/model"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

const (
	insertModQuery = `
insert into mod_info (
    id
	, name
	, url
	, source_url
	, description
	, class_id
	, game_id
	, versions
	, categories
	, authors
) values (
    :id
	, :name
	, :url
	, :source_url
	, :description
	, :class_id
	, :game_id
	, json(:versions)
	, json(:categories)
	, json(:authors)
)
`
)

type ModRepository struct {
	db *sqlx.DB
}

func NewModRepository(db *sqlx.DB) *ModRepository {
	return &ModRepository{db: db}
}

func (r *ModRepository) Save(ctx context.Context, m model.Mod) {
	v, _ := json.Marshal(m.Versions)
	c, _ := json.Marshal(m.Category)
	a, _ := json.Marshal(m.Authors)

	args := map[string]interface{}{
		"id":          m.ID,
		"name":        m.Name,
		"class_id":    m.ClassID,
		"url":         m.URL,
		"source_url":  m.SourceURL,
		"wiki_url":    m.WikiURL,
		"description": m.Description,
		"game_id":     m.GameID,
		"authors":     string(a),
		"versions":    string(v),
		"categories":  string(c),
	}

	slog.With(
		slog.String("authors", string(a)),
		slog.String("versions", string(v)),
		slog.String("category", string(c)),
	).Debug("mod json values")

	_, err := r.db.NamedExecContext(ctx, insertModQuery, args)
	if err != nil {
		err = fmt.Errorf("inserting mod data to local cache: %w", err)
		slog.With("error", err).Warn("ModRepository.Save")
	}
}
