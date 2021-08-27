package main

import (
	"go-booking/api/flight-api/handlers"
	"go-booking/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	flightConn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	flightClient := pb.NewBookingFlightClient(flightConn)

	h := handlers.NewFlightHandler(flightClient)

	g := gin.Default()

	gr := g.Group("/v1/api")
	gr.POST("/create", h.CreateFlight)
	gr.PUT("/update", h.UpdateFlight)
	gr.GET("/search", h.FindFlight)

	g.Run(":3334")
}
