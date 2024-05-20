package store

import (
	"database/sql"
	pb "example/go_crud_grpc/proto"
)

func Create(db *sql.DB, st *pb.Student) (sql.Result, error) {

	result, err := db.Exec(InsertStudentQuery, st.Name, st.StudentId, st.Class, st.Email, st.Address)
	return result, err
}
