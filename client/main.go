package main

import (
	pb "github.com/bostigger/go-grpc-basic/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	if err != nil {
		log.Fatalf("Client could not connect to the server %v", err)
	}

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameLists{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	//callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)

}
