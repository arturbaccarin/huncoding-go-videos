// https://youtu.be/PJr5BT73u38
// 18:55
package main

import (
	"flag"
	"log"
	"raftconsensus/pkg/store"
	"strings"

	"github.com/hashicorp/consul/command/operator/raft"
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

	raftNode := raft.NewNode(config)

	err := raftNode.StartRPCServer(*rpcAddr)
	if err != nil {
		log.Fatal(err)
	}

	raftNode.Start()

	store := store.NewStore(raftNode)

	server := api.NewServer(store, raftNode)

	go func() {
		if err := server.Start(*httpAddr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

}
