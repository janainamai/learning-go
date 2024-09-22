package main

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/janainamai/learning-go/7-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	id := uuid.New().String()
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          id,
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	println("\nListing created categories")
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          id,
		Name:        "New name",
		Description: sql.NullString{String: "New description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	println("\nListing updated categories")
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.DeleteCategory(ctx, id)
	if err != nil {
		panic(err)
	}

	println("\nListing categories after delete")
}
