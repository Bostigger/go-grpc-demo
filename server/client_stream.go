package main

import (
	pb "github.com/bostigger/go-grpc-basic/proto"
	"io"
	"log"
)

func (s *helloServer) sayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageLists{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Message)
		messages = append(messages, "Hello", req.Message)

	}
}
