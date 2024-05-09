package main

import (
	"context"
	"io"
	"log"

	pb "github.com/polyglotdev/grpc-masterclass/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	req := &pb.GreetRequest{FirstName: "Elijah"}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to greet many times: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive a response: %v\n", err)
		}
		log.Printf("Greeting: %s\n", res.Result)
	}
}
