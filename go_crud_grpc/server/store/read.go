package store

import (
	"database/sql"
	pb "example/go_crud_grpc/proto"
)

func Read(db *sql.DB, st *pb.ID) (*pb.Student, error) {

	var result pb.Student

	err := db.QueryRow(ReadStudentQuery, st.Id).Scan(&result.Name, &result.StudentId, &result.Class, &result.Email, &result.Address)

	return &result, err

}
