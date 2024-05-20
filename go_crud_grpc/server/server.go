package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	pb "example/go_crud_grpc/proto"
	"example/go_crud_grpc/server/store"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCRUDServer
}

var db *sql.DB

func main() {

	//Establishing DB connections
	var err error
	db, err = store.EstablishDbConnection()
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
	store.Migrations(db, err)

	// Host grpc service
	listen, err := net.Listen("tcp", os.Getenv("SERVER_ADDR"))
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())

}

func (s *server) CreateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error) {
	// If ID is null, return specific error
	if st.StudentId == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	result, err := store.Create(db, st)
	if err != nil {
		log.Fatal("Yaha pe error aa rha ", err)
	}
	fmt.Println("Insert Success with result :: ", result)

	return &pb.ID{Id: st.StudentId}, err

}

func (s *server) ReadStudent(ctx context.Context, st *pb.ID) (*pb.Student, error) {
	// If ID is null, return specific error

	if st.Id == "" {
		return &pb.Student{}, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	res, err := store.Read(db, st)

	if err != nil {
		log.Printf("Error retrieving employee with id: %s, error: %v", st.Id, err)
		return nil, err
	}
	return res, nil
}

func (s *server) UpdateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error) {
	// If ID is null, return specific error
	if st.StudentId == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}
	err := store.Update(db, st)
	if err != nil {
		log.Fatalf("Error executing SQL statement: %v", err)
	}
	return &pb.ID{Id: st.StudentId}, err
}

func (s *server) DeleteStudent(ctx context.Context, st *pb.ID) (*pb.ID, error) {
	err := store.Delete(db, st)
	return &pb.ID{Id: st.Id}, err
}
