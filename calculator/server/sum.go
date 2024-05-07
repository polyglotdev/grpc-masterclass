package main

import (
	"context"
	"log"

	pb "github.com/polyglotdev/grpc-masterclass/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Received: %v", in)
	result := in.FirstNumber + in.SecondNumber
	return &pb.SumResponse{Result: result}, nil
}
