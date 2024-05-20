package store

import (
	"database/sql"
	pb "example/go_crud_grpc/proto"
	"log"
)

func Delete(db *sql.DB, st *pb.ID) error {

	_, err := db.Exec(DeleteStudentQuery, st.Id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
