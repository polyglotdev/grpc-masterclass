package main

import (
	"context"
	"log"

	pb "github.com/polyglotdev/grpc-masterclass/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	r, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 10, SecondNumber: 20})
	if err != nil {
		log.Fatalf("Failed to sum: %v\n", err)
	}
	log.Printf("Sum: %d", r.Result)
}
