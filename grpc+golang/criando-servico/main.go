package main

/*
gRPC framework open-source de RPC
Utiliza HTTP/2 para transporte e Protocol Buffers (progobufs) para serialização

Protobufs é um mecanismo de serialização de dados estruturados, desenvolvido pelo google

Maneira eficiente e flexível para definir e trocar dados entre serviços

.proto -> usado para definir a estrutura dos dados e as interfaces de serviços.

O problema é se alterar o contrato, pode alterar muitas coisas.

Tem que garantir a ordem e o tipo dos campos (no json só os tipos do campo)

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

// protoc --go_out=. --go-grpc_out=. ./example.proto
*/
