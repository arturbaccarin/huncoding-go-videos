package client

import (
	"clientstreaming/pb"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func Run() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer dial.Close()

	client := pb.NewTemperatureServiceClient(dial)

	stream, err := client.RecordTemperature(context.Background())
	if err != nil {
		panic(err)
	}

	temperatures := []float32{10.0, 20.0, 30.0}

	for _, temperature := range temperatures {
		fmt.Printf("Sending temperature: %f\n", temperature)

		err := stream.Send(&pb.TemperatureRequest{
			Temperature: temperature,
		})
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
	}

	recv, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Temperature received: %f\n", recv.AverageTemperature)
}
