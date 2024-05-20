package store

import (
	"context"
	pb "example/go_crud_grpc/proto"
	"log"
)

func (s *Store) Delete(ctx context.Context, st *pb.ID) error {

	_, err := s.db.Exec(DeleteStudentQuery, st.Id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
