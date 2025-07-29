// https://youtu.be/PJr5BT73u38
// 12:00
package main

// commit of the day - sorry

import (
	"flag"
	"log"
	"strings"

	"github.com/hashicorp/raft"
)

func main() {

	nodeID := flag.String("id", "", "Node ID")
	httpAddr := flag.String("http", ":8080", "HTTP address")
	rpcAddr := flag.String("rpc", ":9090", "RPC address")
	peers := flag.String("peers", "", "Comma-separated list of peer addresses")
	flag.Parse()

	if *nodeID == "" {
		log.Fatal("Node ID is required")
	}

	// Create raft node
	config := raft.DefaultConfig()
	config.ID = *nodeID

	// Parse peers
	if *peers != "" {
		config.Peers = strings.Split(*peers, ",")
	}
}
