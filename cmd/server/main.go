package main

import (
	"api/pkg/adder"
	"api/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := adder.GRPCServer{}
	api.RegisterAdderServer(s, &srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
