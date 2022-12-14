package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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
	//doServerStreaming(ctx, client)

	//client streaming
	//doClientStreaming(ctx, client)

	//bidirectional streaming
	//doBiDirectionalStreaming(ctx, client)

	//timeout
	doRequestResponseWithTimeout(ctx, client)
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

func doClientStreaming(ctx context.Context, client proto.AppServiceClient) {
	data := []int32{3, 1, 4, 2, 5, 8, 6, 7, 9}
	clientStream, err := client.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, val := range data {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Sending %d for average calculation\n", val)
		req := &proto.AverageRequest{
			No: val,
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}
	resp, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average : ", resp.GetResult())
}

func doBiDirectionalStreaming(ctx context.Context, client proto.AppServiceClient) {
	personNames := []proto.PersonName{
		proto.PersonName{FirstName: "Magesh", LastName: "Kuppan"},
		proto.PersonName{FirstName: "Suresh", LastName: "Kannan"},
		proto.PersonName{FirstName: "Rajesh", LastName: "Pandit"},
		proto.PersonName{FirstName: "Ganesh", LastName: "Kumar"},
		proto.PersonName{FirstName: "Ramesh", LastName: "Jayaraman"},
	}
	clientStream, err := client.Greet(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan struct{})
	go func() {
		for {
			resp, err := clientStream.Recv()
			if err != nil {
				fmt.Println(err)
				done <- struct{}{}
				break
			}
			fmt.Println(resp.GetGreetMessage())
		}
	}()
	for _, personName := range personNames {
		time.Sleep(500 * time.Millisecond)
		req := &proto.GreetRequest{
			Person: &personName,
		}
		fmt.Println("Sending : ", personName.FirstName, personName.LastName)
		clientStream.Send(req)
	}
	<-done
}

func doRequestResponseWithTimeout(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	res, err := client.Add(timeoutCtx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout error")
				return
			} else {
				fmt.Println("Unknown error", err)
				return
			}
		}
		log.Fatalln(err)
	}
	fmt.Println(res.GetResult())
	fmt.Println(res.GetDummy())
}
