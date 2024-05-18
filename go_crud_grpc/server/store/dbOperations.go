package dbstore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	pb "example/go_crud_grpc/proto"
	my_queries "example/go_crud_grpc/server/store/queries"

	"github.com/joho/godotenv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func EstablishDbConnection() (*sql.DB, error) {
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatal("Error loading .env file")
	}

	// Open up our database connection.
	// I've set up a database on my local machine using mysql and mysql workbench.
	// The database is called test_db

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", connectionString)
	return db, err

}

func Create(db *sql.DB, st *pb.Student) (sql.Result, error) {
	if st.StudentId == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}
	result, err := db.Exec(my_queries.InsertStudentQuery, st.Name, st.StudentId, st.Class, st.Email, st.Address)
	return result, err
}

func Read(db *sql.DB, st *pb.ID) (*pb.Student, error) {
	if st.Id == "" {
		return &pb.Student{}, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	var result pb.Student

	err := db.QueryRow(my_queries.ReadStudentQuery, st.Id).Scan(&result.Name, &result.StudentId, &result.Class, &result.Email, &result.Address)

	return &result, err

}

func Update(db *sql.DB, st *pb.Student) error {

	stmt, err := db.Prepare(my_queries.UpdateStudentQuery)
	if err != nil {
		log.Fatalf("Error preparing SQL statement: %v", err)
	}

	_, err = stmt.Exec(st.Name, st.Class, st.Email, st.Address, st.StudentId)

	return err
}

func Delete(db *sql.DB, st *pb.ID) error {

	_, err := db.Exec(my_queries.DeleteStudentQuery, st.Id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
