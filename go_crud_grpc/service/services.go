package service

import (
	"context"
	"database/sql"
	pb "example/go_crud_grpc/proto"
	"example/go_crud_grpc/store/student"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	store *store.Store
}

func NewService(db *sql.DB) *Service {
	store := store.New(db)
	return &Service{
		store: store,
	}
}

func (s *Service) CreateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error) {
	// If ID is null, return specific error
	if st.StudentId == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	result, err := s.store.Create(ctx, st)
	if err != nil {
		log.Fatal("Yaha pe error aa rha ", err)
	}
	fmt.Println("Insert Success with result :: ", result)

	return &pb.ID{Id: st.StudentId}, err

}

func (s *Service) ReadStudent(ctx context.Context, st *pb.ID) (*pb.Student, error) {
	// If ID is null, return specific error

	if st.Id == "" {
		return &pb.Student{}, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}

	res, err := s.store.Read(ctx, st)

	if err != nil {
		log.Printf("Error retrieving employee with id: %s, error: %v", st.Id, err)
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error) {
	// If ID is null, return specific error
	if st.StudentId == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}
	err := s.store.Update(ctx, st)
	if err != nil {
		log.Fatalf("Error executing SQL statement: %v", err)
	}
	return &pb.ID{Id: st.StudentId}, err
}

func (s *Service) DeleteStudent(ctx context.Context, st *pb.ID) (*pb.ID, error) {
	err := s.store.Delete(ctx, st)
	return &pb.ID{Id: st.Id}, err
}
