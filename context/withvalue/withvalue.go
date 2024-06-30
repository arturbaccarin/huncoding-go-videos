package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", "value")
	printCtxValue(ctx)
}

func printCtxValue(ctx context.Context) {
	fmt.Println(ctx.Value("key"))
}
