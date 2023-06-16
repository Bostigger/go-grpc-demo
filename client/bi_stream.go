package main

import (
	"context"
	pb "github.com/bostigger/go-grpc-basic/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameLists) {
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		return
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequests{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Printf("Error send the request %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	err = stream.CloseSend()
	if err != nil {
		return
	}
	<-waitc
	log.Printf("Bidirectional Streaming ended")
}
