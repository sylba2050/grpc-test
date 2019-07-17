package main

import (
	"context"
	"fmt"
	pb "github.com/sylba2050/grpc-test/helloworld"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	message := &pb.HelloRequest{Greeting: "tama"}
	res, err := client.SayHello(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
