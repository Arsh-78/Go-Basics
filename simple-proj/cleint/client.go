package main

import (
	"context"
	pb "example/chat/protos"
	"log"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer con.Close()

	c := pb.NewChatServiceClient(con)

	response, err := c.SayHello(context.Background(), &pb.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
