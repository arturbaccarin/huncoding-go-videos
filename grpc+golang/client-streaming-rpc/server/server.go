package server

import (
	"clientstreaming/pb"
	"io"
	"net"

	"google.golang.org/grpc"
)

type TemperatureServer struct {
	pb.UnimplementedTemperatureServiceServer
}

func (*TemperatureServer) RecordTemperature(stream grpc.ClientStreamingServer[pb.TemperatureRequest, pb.TemperatureResponse]) error {
	var sum float32
	var count int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.TemperatureResponse{
				AverageTemperature: sum / float32(count),
			})
		}

		if err != nil {
			return err
		}

		sum += req.Temperature
		count++
	}
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTemperatureServiceServer(grpcServer, &TemperatureServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
