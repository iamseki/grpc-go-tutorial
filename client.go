package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/iamseki/grpc-go-tutorial/chat"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client !",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Reponse from Server: %s", response.Body)
}
