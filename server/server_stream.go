package main

import (
	pb "github.com/bostigger/go-grpc-basic/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameLists, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got the follwoing requests : %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
