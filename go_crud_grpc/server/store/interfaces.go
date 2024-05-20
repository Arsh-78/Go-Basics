package store

import (
	"context"
	pb "example/go_crud_grpc/proto"
)

type store interface {
	Create(ctx context.Context, st *pb.Student)
	Delete(ctx context.Context, st *pb.ID) error
	Read(ctx context.Context, st *pb.ID) (*pb.Student, error)
	Update(ctx context.Context, st *pb.Student) error
}
