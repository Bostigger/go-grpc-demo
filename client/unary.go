package main

import (
	"context"
	pb "github.com/bostigger/go-grpc-basic/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	hello, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		return
	}
	log.Printf("%v", hello.Message)

}
