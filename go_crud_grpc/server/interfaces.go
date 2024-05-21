package main

import (
	"context"
	pb "example/go_crud_grpc/proto"
)

type Server interface {
	CreateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error)
	ReadStudent(ctx context.Context, st *pb.ID) (*pb.Student, error)
	UpdateStudent(ctx context.Context, st *pb.Student) (*pb.ID, error)
	DeleteStudent(ctx context.Context, st *pb.ID) (*pb.ID, error)
}
