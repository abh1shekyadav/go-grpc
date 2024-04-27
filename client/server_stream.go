package main

import (
	"context"
	"io"
	"log"

	pb "github.com/abh1shekyadav/go-grpc/proto"
)


func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Println("Streaming has started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Println("Streaming finished")
}