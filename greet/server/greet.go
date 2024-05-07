package main

import (
	"context"
	"log"

	pb "github.com/polyglotdev/grpc-masterclass/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
