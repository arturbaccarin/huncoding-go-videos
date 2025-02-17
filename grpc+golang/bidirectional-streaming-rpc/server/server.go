package server

import (
	"birpc/pb"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type StockServiceServer struct {
	pb.UnimplementedStockServiceServer
}

func (*StockServiceServer) StreamStockPrices(stream grpc.BidiStreamingServer[pb.StockRequest, pb.StockResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return err
		}

		if err != nil {
			return err
		}

		symbol := req.GetSymbol()
		fmt.Println("Received symbol: ", symbol)

		go func(symbol string) {
			for i := 0; i < 50; i++ {
				price := rand.Float32() * 100
				err := stream.Send(&pb.StockResponse{
					Symbol: symbol,
					Price:  price,
				})
				if err != nil {
					panic(fmt.Sprintln("Error sending stock price: ", err))
				}

				time.Sleep(1 * time.Second)
			}
		}(symbol)
	}
}

func RunServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStockServiceServer(grpcServer, &StockServiceServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
