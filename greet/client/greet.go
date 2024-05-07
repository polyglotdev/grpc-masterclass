package main

import (
	"context"
	"log"

	pb "github.com/polyglotdev/grpc-masterclass/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Elijah"})
	if err != nil {
		log.Fatalf("Failed to greet: %v\n", err)
	}
	log.Printf("Greeting: %s", r.Result)
}
