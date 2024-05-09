package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/polyglotdev/grpc-masterclass/greet/proto"
)

// GreetManyTimes is a function that sends a greeting response multiple times to the client.
// It takes a GreetRequest and a GreetService_GreetManyTimesServer as input parameters.
// The function iterates 10 times, sending a greeting response to the client with a formatted message.
// The formatted message includes the client's first name and a number indicating the iteration.
// The function uses the stream.Send method to send the response to the client.
// After sending each response, the function sleeps for 1000 milliseconds.
// If an error occurs during the sending of the response, the function does not handle it.
// The function returns nil to indicate a successful execution.
//
// Parameters:
//
// req *pb.GreetRequest: The request message sent by the client.
// stream pb.GreetService_GreetManyTimesServer: The server-side stream used to send responses to the client.
//
// Returns:
//
// error: An error if the function encounters an issue while sending the response.
func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", req)

	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s, number %d", req.FirstName, i)
		res := &pb.GreetResponse{
			Result: result,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}
