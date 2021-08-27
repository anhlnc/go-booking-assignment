package main

import (
	"go-booking/api/booking-api/handlers"
	"go-booking/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	bookingConn, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	bookingClient := pb.NewBbookingClient(bookingConn)

	h := handlers.NewBookingHandler(bookingClient)

	g := gin.Default()

	gr := g.Group("/v1/api")
	gr.POST("/create", h.AssignBooking)
	gr.PUT("/delete", h.DeleteBooking)
	gr.GET("/view", h.ViewBooking)

	g.Run(":3335")
}
