package repository

import (
	"context"
	"database/sql"

	"github.com/janainamai/learning-go/8-uow/internal/db"
	"github.com/janainamai/learning-go/8-uow/internal/entity"
)

type (
	CategoryRepositoryInterface interface {
		Insert(ctx context.Context, category entity.Category) error
	}

	CategoryRepository struct {
		DB      *sql.DB
		Queries *db.Queries
	}
)

func NewCategoryRepository(dtb *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
