package main

import (
	"log"
	"net"

	"github.com/iamseki/grpc-go-tutorial/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 : %v", err)
	}

	s := chat.Server{}

	gsrv := grpc.NewServer()

	chat.RegisterChatServiceServer(gsrv, &s)

	if err := gsrv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server on port 9000 : %v", err)
	}

}
