package main

import (
	"log"

	pb "github.com/polyglotdev/grpc-masterclass/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Received Primes request: %v", req)

	number := req.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			if err := stream.Send(&pb.PrimeResponse{Result: divisor}); err != nil {
				return err
			}

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
