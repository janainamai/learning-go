package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/janainamai/learning-go/5-grpc/internal/database"
	"github.com/janainamai/learning-go/5-grpc/internal/pb"
	"github.com/janainamai/learning-go/5-grpc/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// init database connection
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// init resources
	categoryDB := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDB)

	// init server
	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryService)
	reflection.Register(server) // para trabalhar com Evans

	// config tcp port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on 50051")
	// execute server
	if err := server.Serve(lis); err != nil {
		panic(err)
	}

}
