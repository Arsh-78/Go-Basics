package store

import (
	"context"
	"database/sql"
	pb "example/go_crud_grpc/proto"
)

func (s *Store) Create(ctx context.Context, st *pb.Student) (sql.Result, error) {

	result, err := s.db.Exec(InsertStudentQuery, st.Name, st.StudentId, st.Class, st.Email, st.Address)
	return result, err
}
