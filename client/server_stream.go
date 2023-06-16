package main

import (
	"context"
	pb "github.com/bostigger/go-grpc-basic/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameLists) {
	log.Printf("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("coudnt send name %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")

}
