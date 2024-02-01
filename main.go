package main

import (
	"context"
	"log"
	"net"
	userpb "test/test/protobuf/user"
	wearablepb "test/test/protobuf/wearable"
	servertypes "test/servertypes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

)

type userservice struct {
	userpb.UnimplementedUserServiceServer
}

func (u *userservice) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {

	cp := &userpb.CPU{
		Name:  "monu",
		Brand: req.Id,
		Cores: 8,
	}

	res := &userpb.GetUserResponse{
		Cpu: cp,
	}

	return res, nil
}

func main() {
	log.Println("hjfghjdfg")
	lis, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		log.Println("error in setting up tcp server",err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userservice{})
	wearablepb.RegisterWearableServiceServer(grpcServer, &servertypes.Wearable{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

}
