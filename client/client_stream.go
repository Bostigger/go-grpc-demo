package main

import (
	"context"
	pb "github.com/bostigger/go-grpc-basic/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameLists) {
	log.Printf("Client Streaming Started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		return
	}

	for _, name := range names.Names {
		req := &pb.HelloResponse{
			Message: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with message %s", name)
		time.Sleep(2 * time.Second)

		res, err := stream.CloseAndRecv()
		log.Printf("Client streaming finished")
		if err != nil {
			log.Fatalf("Error ")
		}
		log.Printf("%v", res.Messages)

	}
}
