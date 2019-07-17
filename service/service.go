package service

import (
	"context"
	"errors"
	pb "github.com/sylba2050/grpc-test/helloworld"
)

type MyService struct {
}

func (s *MyService) SayHello(ctx context.Context, message *pb.HelloRequest) (*pb.HelloResponse, error) {
	switch message.Greeting {
	case "tama":
		return &pb.HelloResponse{
			Reply: "タマ",
		}, nil
	case "mike":
		return &pb.HelloResponse{
			Reply: "ミケ",
		}, nil

	}
	return nil, errors.New("Not Found YourCat")
}
