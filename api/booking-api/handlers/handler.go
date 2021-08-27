package handlers

import (
	"go-booking/api/booking-api/requests"
	"go-booking/api/booking-api/responses"
	"go-booking/pb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookingHandler interface {
	AssignBooking(c *gin.Context)
	ViewBooking(c *gin.Context)
	DeleteBooking(c *gin.Context)
}

type bookingHandler struct {
	bookingClient pb.BbookingClient
}

func (h *bookingHandler) AssignBooking(c *gin.Context) {
	req := requests.AssignBookingRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errMessage := make([]string, 0)
			for _, v := range validateErr {
				errMessage = append(errMessage, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessage,
			})

			return
		}
	}
	pReq := &pb.AssignBookingRequest{
		CustomerId: req.Customer_id,
		FlightId:   req.Flight_id,
		Status:     req.Status,
		Code:       req.Code,
		BookDate:   time.Now().String(),
	}

	pRes, err := h.bookingClient.AssignBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.BookingResponse{
		Code:         pRes.Code,
		Booking_Date: pRes.BookDate,
		Customer_id:  pRes.CustomerId,
		Flight_id:    pRes.FlightId,
		Status:       pRes.Status,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *bookingHandler) ViewBooking(c *gin.Context) {
	req := requests.ViewBookingRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}
	pReq := &pb.ViewBookingRequest{
		Code:       req.Code,
		CustomerId: req.Customer_id,
	}
	pRes, err := h.bookingClient.ViewBookings(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pRes,
	})
}

func (h *bookingHandler) DeleteBooking(c *gin.Context) {
	req := requests.DeleteBookingRequest{}
	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}
	pReq := &pb.DeleteBookingRequest{
		Id: req.Id,
	}
	bReq := &pb.GetBookingByIDRequest{
		Id: req.Id,
	}
	booked, errGet := h.bookingClient.GetBookingByID(c.Request.Context(), bReq)
	if errGet != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  errGet.Error(),
		})
		return
	}
	_, err := h.bookingClient.DeleteBooking(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := responses.DeleteBookingResponse{
		Code:        booked.Code,
		Cancel_date: time.Now(),
		Customer_id: booked.CustomerId,
		Flight_id:   booked.FlightId,
		Status:      "Cancel",
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func NewBookingHandler(bookingClient pb.BbookingClient) BookingHandler {
	return &bookingHandler{
		bookingClient: bookingClient,
	}
}
