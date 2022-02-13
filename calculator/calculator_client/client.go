package main

import (
	"context"
	"fmt"
	"go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
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
	//doServerStreaming(c)
	//doClientStreaming(c)
	// doBiDiStreaming(c)

	doErrorUnary(c)
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
		X: 1277124,
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	requests := []*calculatorpb.CalculatorStreamingRequest{
		&calculatorpb.CalculatorStreamingRequest{
			X: int32(1),
		},
		&calculatorpb.CalculatorStreamingRequest{
			X: int32(2),
		},
		&calculatorpb.CalculatorStreamingRequest{
			X: int32(3),
		},
		&calculatorpb.CalculatorStreamingRequest{
			X: int32(4),
		},
	}

	stream, err := c.CalculateAverage(context.Background())
	if err != nil {
		log.Fatalf("Error streaming data: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Uh oh, spaghetti-o: %v", err)
	}
	fmt.Printf("Average is: %v\n", resp)
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do BiDi streaming")
	stream, err := c.CalculateStreamingMax(context.Background())
	if err != nil {
		log.Fatalf("Error starting to stream to server: %v", err)
	}

	waitc := make(chan struct{})
	requests := []int32{1, 5, 3, 6, 2, 20}

	go func() {
		for _, i := range requests {
			req := &calculatorpb.CalculatorStreamingRequest{
				X: i,
			}
			fmt.Println("Sending value of %v", req.GetX())
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Closing connection due to %v", err)
				break
			}
			fmt.Println("New max of %v\n", res.GetX())
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")
	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	resp, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v\n", err)
			return
		}
	}
	fmt.Printf("Result of square root of %v is %v\n", n, resp.GetNumberRoot())
}
