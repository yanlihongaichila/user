package main

import (
	"context"
	"flag"
	pb "gateway/proto/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUserServer(s, new(AAA))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type AAA struct {
	pb.UnimplementedUserServer
}

func (a *AAA) Ping(context.Context, *pb.Request) (*pb.Response, error) {
	return &pb.Response{Pong: "111"}, nil
}
