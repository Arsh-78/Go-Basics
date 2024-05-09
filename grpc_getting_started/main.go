package main

import (
	"fmt"
	"log"
	"net"

	protos "github.com/truesch/grpc_getting_started/protos/translation"
	"github.com/truesch/grpc_getting_started/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()

	trans := server.NewTranslation()

	reflection.Register(s)

	protos.RegisterTranslationServer(s, trans)

	tl, err := net.Listen("tcp", "localhost:8765")

	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}

	s.Serve()

}
