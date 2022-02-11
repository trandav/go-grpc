package main

import (
	"context"
	"fmt"
	"go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello, I'm a client.")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %v", c)

	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &calculatorpb.CalculatorRequest{
		X: 3,
		Y: 10,
	}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Geret RPC: %v", err)
	}
	log.Printf("Response from Calculate: %v", res.Sum)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculatorStreamingRequest{
		X: 120,
	}
	resStream, err := c.CalculatePrimeStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("Error with the streaming request %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error with the message %v", err)
		}

		log.Printf("Prime number %v", msg.X)
	}
}
