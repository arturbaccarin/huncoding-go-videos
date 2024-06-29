package main

import (
	"context"
	"fmt"
	"goroutineandasync/mongodb"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func BenchmarkAsyncFlow(b *testing.B) {

	ctx := context.Background()

	user := mongodb.User{
		Name: "Artur",
		Age:  30,
	}

	for i := 0; i < b.N; i++ {
		chanUser := make(chan mongodb.ChanUser)
		go mongodb.InsertUserAsyncMongoDB(ctx, user, chanUser)

		chanUserValue := <-chanUser
		if chanUserValue.Err != nil {
			b.FailNow()
			return
		}

		fmt.Println(chanUserValue.User)
	}
}
