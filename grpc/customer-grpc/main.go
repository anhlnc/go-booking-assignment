package main

import (
	"fmt"
	"go-booking/grpc/customer-grpc/handlers"
	"go-booking/grpc/customer-grpc/repositories"
	"go-booking/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	BookingConn, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	bookingClient := pb.NewBbookingClient(BookingConn)
	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	customerRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewCustomerHandler(bookingClient, customerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterBookingCustomerServer(s, h)

	fmt.Println("Listen at port: 2222")

	s.Serve(listen)
}
