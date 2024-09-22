package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/janainamai/learning-go/7-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type (
	CourseDB struct {
		dbConn *sql.DB
		*db.Queries
	}

	CourseParams struct {
		ID          string
		Name        string
		Description sql.NullString
		Price       float64
	}

	CategoryParams struct {
		ID          string
		Name        string
		Description sql.NullString
	}
)

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := db.New(tx)
	err = fn(queries) // executa função recebida com transação
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %v", errRb, err)
		}

		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			CategoryID:  argsCategory.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	courseParams := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Go course", Valid: true},
		Price:       10.95,
	}
	categoryParams := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "TI",
		Description: sql.NullString{String: "Tecnologia da informação", Valid: true},
	}

	courseDB := NewCourseDB(dbConn)
	err = courseDB.CreateCourseAndCategory(ctx, categoryParams, courseParams)
	if err != nil {
		panic(err)
	}

	queries := db.New(dbConn)
	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("Category name: %s, Course ID: %s, Name: %s, Description: %s",
			course.CategoryName, course.ID, course.Name, course.Description.String)
	}

}
