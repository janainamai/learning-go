package repository

import (
	"context"
	"database/sql"

	"github.com/janainamai/learning-go/8-uow/internal/db"
	"github.com/janainamai/learning-go/8-uow/internal/entity"
)

type (
	CourseRepositoryInterface interface {
		Insert(ctx context.Context, course entity.Course) error
	}

	CourseRepository struct {
		DB      *sql.DB
		Queries *db.Queries
	}
)

func NewCourseRepository(dtb *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: int32(course.CategoryID),
	})
}
