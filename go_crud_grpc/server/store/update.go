package store

import (
	"database/sql"
	pb "example/go_crud_grpc/proto"
	"log"
)

func Update(db *sql.DB, st *pb.Student) error {
	stmt, err := db.Prepare(UpdateStudentQuery)
	if err != nil {
		log.Fatalf("Error preparing SQL statement: %v", err)
	}
	_, err = stmt.Exec(st.Name, st.Class, st.Email, st.Address, st.StudentId)
	return err
}
