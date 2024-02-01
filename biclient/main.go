package main

import (
	"context"
	"fmt"
	"io"
	"log"
	wearablepb "test/test/protobuf/wearable"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9000", opts...)
	if err != nil {
		log.Println("fail to dial:", err)
	}
	defer conn.Close()
	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.CalculateBeatsPerMinute(context.Background())

	if err != nil {
		log.Println("opening", err)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("here inb the lop")
		 err := stream.Send(&wearablepb.CalculateRequest{
			Id:     "tejas",
			Value:  uint32(i),
			Minute: uint32(i + 1),
		}); 
		if err != nil {
			fmt.Println("sendig error", err)
		}

	}
	if err:=stream.CloseSend(); err !=nil {
		log.Println("closing errror",err)
	}
	for {
		res,err := stream.Recv()
		if err == io.EOF {
			log.Println("end of recieving files")
			break
		}
		fmt.Println("Average value is ", res.GetAverage())
	}
}
