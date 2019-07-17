package main

import (
	pb "github.com/sylba2050/grpc-test/helloworld"
	"github.com/sylba2050/grpc-test/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	s := &service.MyService{}
	pb.RegisterHelloServiceServer(server, s)
	server.Serve(listenPort)
}
