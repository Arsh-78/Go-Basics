package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "example/go_crud_grpc/proto"

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

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	var err error
	connectionString := "root:new-password@tcp(127.0.0.1:3306)/test_db"
	db, err = sql.Open("mysql", connectionString)
	fmt.Println("db in main", db)

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
	fmt.Println("Successfully connected to the database!")

	// Host grpc service
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
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
	fmt.Println("db in createstud ", db)
	result, err := db.Exec(`INSERT INTO studentTwo (name, studentId, class, email, address) VALUES (?, ?, ?, ?, ?)`, st.Name, st.StudentId, st.Class, st.Email, st.Address)
	if err != nil {
		log.Fatal("Yaha pe error aa rha ", err)
	}
	fmt.Println("Insert Success with result :: ", result)

	return &pb.ID{Id: st.StudentId}, err

}

func (s *server) ReadStudent(ctx context.Context, st *pb.ID) (*pb.Student, error) {
	// If ID is null, return specific error
	if st.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	var result pb.Student

	query := "SELECT * FROM studentTwo WHERE studentId = ?"
	err := db.QueryRow(query, st.Id).Scan(&result.Name, &result.StudentId, &result.Class, &result.Email, &result.Address)

	if err != nil {
		log.Printf("Error retrieving employee with id: %s, error: %v", st.Id, err)
		return nil, err
	}

	return &result, nil
}

func (s *server) UpdateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error) {
	// If ID is null, return specific error
	if st.StudentId == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	stmt, err := db.Prepare("UPDATE studentTwo SET name = ?, class = ? , email = ? ,address = ? WHERE studentId = ?")
	if err != nil {
		log.Fatalf("Error preparing SQL statement: %v", err)
	}

	_, err = stmt.Exec(st.Name, st.Class, st.Email, st.Address, st.StudentId)

	if err != nil {
		log.Fatalf("Error executing SQL statement: %v", err)
	}

	return &pb.ID{Id: st.StudentId}, err
}

func (s *server) DeleteStudent(ctx context.Context, st *pb.ID) (*pb.ID, error) {
	_, err := db.Exec(`DELETE FROM studentTwo WHERE studentId = ?`, st.Id)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.ID{Id: st.Id}, err
}
