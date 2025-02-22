// https://youtu.be/1GSPmP3NSxA
package main

import (
	"unaryrpc/client"
	"unaryrpc/server"
)

// protoc --go_out=. --go-grpc_out=. user.proto
func main() {
	go server.Run()
	client.Run()
}

// openssl req -x509 -newkey rsa:2048 -keyout server.key -out server.crt -days 365 -nodes -config openssl-san.cnf -extensions v3_req
