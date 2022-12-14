package main

import (
	"context"
	"errors"
	"fmt"
	"grpc-demo/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type appServerImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *appServerImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	fmt.Println("Service processing the request......")
	processTimeCh := time.After(5 * time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel request received from the client. Discarding the processing...")
			return nil, errors.New("client cancelled the request")
		case <-processTimeCh:
			res := &proto.AddResponse{
				Result: result,
				Dummy:  []int32{10, 20, 30, 40},
			}
			fmt.Printf("Processing %d and %d & returning result %d\n", x, y, result)
			return res, nil
		}
	}
	return nil, nil
}

func (asi *appServerImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Sending Prime no : %d\n", no)
			resp := &proto.PrimeResponse{
				PrimeNo: no,
			}
			serverStream.Send(resp)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (asi *appServerImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {

	sum := int32(0)
	count := int32(0)
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		sum += req.GetNo()
		count++
	}
	avg := sum / count
	resp := &proto.AverageResponse{
		Result: avg,
	}
	err := serverStream.SendAndClose(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

//bidirectional streaming
func (asi *appServerImpl) Greet(serverStream proto.AppService_GreetServer) error {
	for {
		req, err := serverStream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		person := req.GetPerson()
		msg := fmt.Sprintf("Hi %s %s!", person.GetFirstName(), person.GetLastName())
		resp := &proto.GreetResponse{
			GreetMessage: msg,
		}
		time.Sleep(2 * time.Second)
		e := serverStream.Send(resp)
		if e != nil {
			log.Fatalln(err)
		}
	}
	return nil
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
