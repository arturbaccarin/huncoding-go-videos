package client

import (
	"birpc/pb"
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func RunClient() {
	dial, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer dial.Close()

	client := pb.NewStockServiceClient(dial)

	stream, err := client.StreamStockPrices(context.Background())
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				panic(err)
			}

			log.Printf("Received Stock Price: %v", response.GetPrice())
		}

		close(done)
	}()

	symbols := []string{"MSFT", "GOOG", "AAPL"}
	for _, symbol := range symbols {
		log.Printf("Sending symbol: %v", symbol)
		if err := stream.Send(&pb.StockRequest{
			Symbol: symbol,
		}); err != nil {
			panic(err)
		}
		time.Sleep(2 * time.Second)
	}

	<-done
}
