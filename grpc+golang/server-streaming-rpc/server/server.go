package server

import (
	"fmt"
	"net"
	"serverstreaming/pb"
	"time"

	"google.golang.org/grpc"
)

type StatusServer struct {
	pb.UnimplementedStatusServiceServer
}

func (s *StatusServer) StreamStatus(req *pb.StreamRequest, stream grpc.ServerStreamingServer[pb.StreamResponse]) error {
	taskId := req.TaskId
	fmt.Println("Streaming status for tasj id: ", taskId)

	for i := 0; i < 100; i += 10 {
		status := pb.StreamResponse{
			Message:  "Progressing",
			Progress: int64(i),
		}

		err := stream.Send(&status)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStatusServiceServer(grpcServer, &StatusServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
