package main

import (
	"go-booking/api/customer-api/handlers"
	"go-booking/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	customerConn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	customerClient := pb.NewBookingCustomerClient(customerConn)

	h := handlers.NewCustomerHandler(customerClient)

	g := gin.Default()

	gr := g.Group("/v1/api")
	gr.POST("/create", h.CreateCustomer)
	gr.PUT("/update", h.UpdateCustomer)
	gr.GET("/history", h.BookingHistory)
	gr.PUT("/changePassword", h.ChangePassword)

	g.Run(":3333")
}
