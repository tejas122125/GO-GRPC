package main

import (
	"context"
	"fmt"
	"log"
	userpb "test/test/protobuf/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



func main (){
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9000", opts...)
	if err != nil {
		log.Println("fail to dial:", err)
	}

	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Id: "1-2-3",
	})
	if err != nil {
		log.Fatalf("fail to GetUser: %v", err)
	}

	fmt.Printf("%+v\n", res)

}