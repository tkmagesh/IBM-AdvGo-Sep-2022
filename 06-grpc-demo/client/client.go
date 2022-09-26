package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	//request & response
	//doRequestResponse(ctx, client)

	//server streaming
	doServerStreaming(ctx, client)
}

func doRequestResponse(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.GetResult())
	fmt.Println(res.GetDummy())
}

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	primeRequest := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	clientStream, err := client.GeneratePrimes(ctx, primeRequest)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		resp, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("Thats all folks")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Prime Number received : ", resp.GetPrimeNo())
	}
}
