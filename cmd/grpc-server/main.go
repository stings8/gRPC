package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stings8/gRPC/internal/database"
	"github.com/stings8/gRPC/internal/pb"
	"github.com/stings8/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
