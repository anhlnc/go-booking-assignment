package main

import (
	"fmt"
	"go-booking/grpc/flight-grpc/handlers"
	"go-booking/grpc/flight-grpc/repositories"
	"go-booking/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	flightRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewFlightHandler(flightRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterBookingFlightServer(s, h)

	fmt.Println("Listen at port: 2223")

	s.Serve(listen)
}
