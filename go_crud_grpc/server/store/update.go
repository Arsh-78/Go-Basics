package store

import (
	"context"
	pb "example/go_crud_grpc/proto"
	"log"
)

func (s *Store) Update(ctx context.Context, st *pb.Student) error {
	stmt, err := s.db.Prepare(UpdateStudentQuery)
	if err != nil {
		log.Fatalf("Error preparing SQL statement: %v", err)
	}
	_, err = stmt.Exec(st.Name, st.Class, st.Email, st.Address, st.StudentId)
	return err
}
