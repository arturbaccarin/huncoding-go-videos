package main

import (
	"flag"
	"fmt"
	"os"
)

// go build .
// ./cli-project --channelName=artur
// ./cli-project --help
func main() {
	channel := flag.String("channelName", "life of boris", "Nome do seu canal")
	flag.Parse()

	os.Mkdir(fmt.Sprintf("./%s", *channel), 077)

	fmt.Println(*channel)
}
