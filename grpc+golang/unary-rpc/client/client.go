package client

import (
	"context"
	"fmt"
	"unaryrpc/pb"

	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer dial.Close()

	userClient := pb.NewUserClient(dial)

	user, err := userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "1",
		Age:  10,
		Name: "John",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("First user created: %v\n", user)

	user, err = userClient.AddUser(context.Background(), &pb.AddUserRequest{
		Id:   "2",
		Age:  20,
		Name: "Jane",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Second user created: %v\n", user)

	getUserResponse, err := userClient.GetUser(context.Background(), &pb.GetUserRequest{
		Id: "1",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("User returned from GetUser method: %v\n", getUserResponse)
}
