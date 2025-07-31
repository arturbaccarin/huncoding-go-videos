package raft

import (
	"log"
	"net"
	"net/rpc"
)

func (n *Node) StartRPCServer(addr string) error {
	server := &RPCServer{node: n}
	rpc.Register(server)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Printf("Failed to accept connection: %v", err)
				continue
			}
			go rpc.ServeConn(conn)
		}
	}()

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
