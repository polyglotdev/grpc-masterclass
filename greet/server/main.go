package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/polyglotdev/grpc-masterclass/greet/proto"
)

var addr = "0.0.0.0:50052"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v\n", addr)

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
