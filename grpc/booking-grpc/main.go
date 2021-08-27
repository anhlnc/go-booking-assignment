package main

import (
	"fmt"
	"go-booking/grpc/booking-grpc/handlers"
	"go-booking/grpc/booking-grpc/repositories"
	"go-booking/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	customerConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	flightConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	customerClient := pb.NewBookingCustomerClient(customerConn)
	flightClient := pb.NewBookingFlightClient(flightConn)

	listen, err := net.Listen("tcp", ":2224")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	bookingRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewBookingHandler(customerClient, flightClient, bookingRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterBbookingServer(s, h)

	fmt.Println("Listen at port: 2224")

	s.Serve(listen)
}
