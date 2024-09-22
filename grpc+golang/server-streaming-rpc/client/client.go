package client

import (
	"context"
	"fmt"
	"io"
	"serverstreaming/pb"

	"google.golang.org/grpc"
)

func Run() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := pb.NewStatusServiceClient(conn)

	stream, err := client.StreamStatus(context.Background(), &pb.StreamRequest{
		TaskId: "1",
	})
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received status: %s, progress: %d%%\n", res.GetMessage(), res.GetProgress())
	}
}
