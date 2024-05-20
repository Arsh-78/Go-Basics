package store

import (
	"context"
	pb "example/go_crud_grpc/proto"
)

func (s *Store) Read(ctx context.Context, st *pb.ID) (*pb.Student, error) {

	var result pb.Student

	err := s.db.QueryRow(ReadStudentQuery, st.Id).Scan(&result.Name, &result.StudentId, &result.Class, &result.Email, &result.Address)

	return &result, err

}
