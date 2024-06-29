package main

import (
	"context"
	"fmt"
	"goroutineandasync/mongodb"
)

func main() {
	err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		panic(err)
	}

	chanUser := make(chan mongodb.ChanUser)

	go mongodb.InsertUserAsyncMongoDB(context.Background(), mongodb.User{
		Name: "Artur",
		Age:  30,
	}, chanUser)

	chanUserValue := <-chanUser
	if chanUserValue.Err != nil {
		panic(chanUserValue.Err)
	}

	fmt.Println(chanUserValue.User)
}
