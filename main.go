package main

import (
	"context"
	"log"
	"net"
	userpb "test/test/protobuf/user"

	"google.golang.org/grpc"
)

type userservice struct{
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
		log.Println("error in setting up tcp server")
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userservice{})
	grpcServer.Serve(lis)

	// type UserServiceServer interface {
	// 	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// 	mustEmbedUnimplementedUserServiceServer()
	// }
}
