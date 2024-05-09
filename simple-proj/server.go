package main

import (
	"context"
	"log"
	"net"

	pb "example/chat/protos"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *pb.Message) (*pb.Message, error) {
	log.Printf("Receive message body from client: %s", req.Body)
	return &pb.Message{Body: "Hello From the Server!"}, nil

}

func main() {

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterChatServiceServer(s, &Server{})
	log.Printf("gRPC server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
