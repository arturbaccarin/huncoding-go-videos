package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"unaryrpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func Run() {
	creds := credentials.NewTLS(&tls.Config{})

	dial, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	defer dial.Close()

	userClient := pb.NewUserClient(dial)

	login, err := userClient.Login(context.Background(), &pb.LoginRequest{
		Username: "admin",
		Password: "admin",
	})

	md := metadata.New(map[string]string{
		"Authorization": "Bearer " + login.Token,
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	user, err := userClient.AddUser(ctx, &pb.AddUserRequest{
		Id:   "1",
		Age:  10,
		Name: "John",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("First user created: %v\n", user)

	user, err = userClient.AddUser(ctx, &pb.AddUserRequest{
		Id:   "2",
		Age:  20,
		Name: "Jane",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Second user created: %v\n", user)

	getUserResponse, err := userClient.GetUser(ctx, &pb.GetUserRequest{
		Id: "1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("User returned from GetUser method: %v\n", getUserResponse)
}
