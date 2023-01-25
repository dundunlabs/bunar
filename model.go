package bunar

import (
	"context"

	"github.com/uptrace/bun"
)

type Model[T any] struct {
	records []T
	sq      *bun.SelectQuery
}

func NewModel[T any](db *bun.DB, model T) *Model[T] {
	records := []T{}

	return &Model[T]{
		records: records,
		sq:      db.NewSelect().Model(&records),
	}
}

func (m *Model[T]) All(ctx context.Context) ([]T, error) {
	err := m.sq.Scan(ctx)
	return m.records, err
}
