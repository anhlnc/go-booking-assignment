package handlers

import (
	"go-booking/api/flight-api/requests"
	"go-booking/api/flight-api/responses"
	"go-booking/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FlightHandler interface {
	CreateFlight(c *gin.Context)
	UpdateFlight(c *gin.Context)
	FindFlight(c *gin.Context)
}

type flightHandler struct {
	flightClient pb.BookingFlightClient
}

func (h *flightHandler) CreateFlight(c *gin.Context) {
	req := requests.CreateFlightRequest{}
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

	pReq := &pb.Flight{

		Name: req.Name,
	}

	pRes, err := h.flightClient.CreateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		ID:             pRes.Id,
		Name:           pRes.Name,
		From:           pReq.From,
		To:             pReq.To,
		Date:           pReq.Date,
		Status:         pReq.Status,
		Available_slot: int(pReq.AvailableSlot),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *flightHandler) UpdateFlight(c *gin.Context) {
	req := requests.UpdateFlightRequest{}
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

	pReq := &pb.Flight{
		Id:            req.Id,
		Name:          req.Name,
		From:          req.From,
		To:            req.To,
		Date:          req.Date,
		Status:        req.Status,
		AvailableSlot: int64(req.Available_slot),
	}

	pRes, err := h.flightClient.UpdateFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		ID:             pRes.Id,
		Name:           pRes.Name,
		From:           pReq.From,
		To:             pReq.To,
		Date:           pReq.Date,
		Status:         pReq.Status,
		Available_slot: int(pReq.AvailableSlot),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *flightHandler) FindFlight(c *gin.Context) {
	req := requests.FindFlightRequest{}
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

	pReq := &pb.FindFlightRequest{
		Name: req.Name,
		From: req.From,
		To:   req.To,
		Date: req.Date,
	}
	pRes, err := h.flightClient.FindFlight(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.FlightResponse{
		ID:             pRes.Id,
		Name:           pRes.Name,
		From:           pRes.From,
		To:             pRes.To,
		Date:           pRes.Date,
		Status:         pRes.Status,
		Available_slot: int(pRes.AvailableSlot),
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func NewFlightHandler(flightClient pb.BookingFlightClient) FlightHandler {
	return &flightHandler{
		flightClient: flightClient,
	}
}
