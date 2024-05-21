package main

import (
	"context"
	mg "example/go_crud_grpc/dbMigrations"
	pb "example/go_crud_grpc/proto"
	ser "example/go_crud_grpc/service"
	"example/go_crud_grpc/store/student"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"log"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedCRUDServer
	service *ser.Service
}

func main() {

	//Establishing DB connections

	db, err := store.EstablishDbConnection()
	fmt.Println("db in main", db)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	//Pinging for connection check
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
	fmt.Println("Successfully connected to the database!")

	//Performing database migrations for table and db if not exists
	mg.Migrations(db, err)

	// Host grpc service
	listen, err := net.Listen("tcp", os.Getenv("SERVER_ADDR"))
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &server{service: ser.NewService(db)})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())

}

func (s *server) CreateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error) {
	// If ID is null, return specific error
	return s.service.CreateStudent(ctx, st)

}

func (s *server) ReadStudent(ctx context.Context, st *pb.ID) (*pb.Student, error) {
	// If ID is null, return specific error
	return s.service.ReadStudent(ctx, st)
}

func (s *server) UpdateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error) {
	// If ID is null, return specific error
	return s.service.UpdateStudent(ctx, st)
}

func (s *server) DeleteStudent(ctx context.Context, st *pb.ID) (*pb.ID, error) {

	return s.service.DeleteStudent(ctx, st)
}
