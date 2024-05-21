package store

import (
	"context"
	"database/sql"
	pb "example/go_crud_grpc/proto"
	"log"
)

func (s *Store) Create(ctx context.Context, st *pb.Student) (sql.Result, error) {

	result, err := s.db.Exec(InsertStudentQuery, st.Name, st.StudentId, st.Class, st.Email, st.Address)
	return result, err
}
func (s *Store) Read(ctx context.Context, st *pb.ID) (*pb.Student, error) {

	var result pb.Student

	err := s.db.QueryRow(ReadStudentQuery, st.Id).Scan(&result.Name, &result.StudentId, &result.Class, &result.Email, &result.Address)

	return &result, err

}

func (s *Store) Update(ctx context.Context, st *pb.Student) (sql.Result, error) {
	// query := UpdateStudentQuery
	// res, err := s.db.ExecContext(ctx, query, st.Name, st.Class, st.Email, st.Address, st.StudentId)
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil

	result, err := s.db.Exec(UpdateStudentQuery, st.Name, st.Class, st.Email, st.Address, st.StudentId)
	if err != nil {
		log.Fatal("failed to execute update query:", err)
	}

	return result, err
}

func (s *Store) Delete(ctx context.Context, st *pb.ID) error {

	_, err := s.db.Exec(DeleteStudentQuery, st.Id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
