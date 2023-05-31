package main

import (
	"api/pkg/api"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("non enough argument")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAdderClient(conn)

	res, err := c.Add(context.Background(), &api.AddRequest{
		X: int32(x),
		Y: int32(y),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Add argument result", res.GetResult())

	resN, err := c.SetName(context.Background(), &api.Person{
		Name:    flag.Arg(2),
		SubName: flag.Arg(3),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(resN.GetFullName())

}
