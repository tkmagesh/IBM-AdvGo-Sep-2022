package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type appServerImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *appServerImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	fmt.Printf("Processing %d and %d & returning result %d\n", x, y, result)
	return res, nil
}

func main() {
	asi := &appServerImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
