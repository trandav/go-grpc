package main

import (
	"context"
	"fmt"
	"go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
}

func (*server) Calculate(ctx context.Context, r *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculate the following: %v + %v", r.X, r.Y)
	result := r.X + r.Y
	rsp := calculatorpb.CalculatorResponse{
		Sum: result,
	}
	return &rsp, nil
}

func (*server) CalculatePrimeStreaming(r *calculatorpb.CalculatorStreamingRequest, stream calculatorpb.CalculatorService_CalculatePrimeStreamingServer) error {
	k := 2
	N := int(r.GetX())
	for N != 1 {
		if N%k == 0 {
			fmt.Printf("This is a factor: %v\n", k)
			N = N / k
			stream.Send(&calculatorpb.CalculatorStreamingResponse{
				X: int32(k),
			})
		} else {
			k = k + 1
		}
	}
	return nil
}

func (*server) CalculateAverage(stream calculatorpb.CalculatorService_CalculateAverageServer) error {
	fmt.Println("Getting a streaming client request.")
	sum := 0
	count := 0
	for {
		x, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calculatorpb.CalculatorAverageResponse{
				X: average,
			})
		}
		if err != nil {
			log.Fatalf("Error with receiving client data: %v", err)
		}
		sum += int(x.GetX())
		count++
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
