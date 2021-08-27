package handlers

import (
	"go-booking/api/customer-api/requests"
	"go-booking/api/customer-api/responses"
	"go-booking/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type CustomerHandler interface {
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	BookingHistory(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type customerHandler struct {
	customerClient pb.BookingCustomerClient
}

func (h *customerHandler) CreateCustomer(c *gin.Context) {
	req := requests.CreateCustomerRequest{}
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
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	pReq := &pb.Customer{
		Name:        req.Name,
		Address:     req.Address,
		LicenseId:   req.License_id,
		PhoneNumber: req.Phone_number,
		Email:       req.Email,
		Password:    string(hash),
		Active:      true,
	}

	pRes, err := h.customerClient.CreateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		ID:           pRes.Id,
		Name:         pRes.Name,
		Address:      pRes.Address,
		License_id:   pRes.LicenseId,
		Phone_number: pRes.PhoneNumber,
		Email:        pRes.Email,
		Password:     pRes.Password,
		Active:       pRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *customerHandler) UpdateCustomer(c *gin.Context) {
	req := requests.UpdateCustomerRequest{}
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
	passHash := req.Password
	if req.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		passHash = string(hash)
	}
	pReq := &pb.Customer{
		Id:          req.Id,
		Name:        req.Name,
		Address:     req.Address,
		LicenseId:   req.License_id,
		PhoneNumber: req.Phone_number,
		Email:       req.Email,
		Password:    passHash,
		Active:      true,
	}

	pRes, err := h.customerClient.UpdateCustomer(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		ID:           pRes.Id,
		Name:         pRes.Name,
		Address:      pRes.Address,
		License_id:   pRes.LicenseId,
		Phone_number: pRes.PhoneNumber,
		Email:        pRes.Email,
		Password:     pRes.Password,
		Active:       pRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
func (h *customerHandler) BookingHistory(c *gin.Context) {
	req := requests.BookingHistoryRequest{}
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
	pReq := &pb.HistoryBookingRequest{
		Id: req.Id,
	}

	pRes, err := h.customerClient.HistoryBooking(c.Request.Context(), pReq)
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

func (h *customerHandler) ChangePassword(c *gin.Context) {
	req := requests.ChangePasswordRequest{}
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

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	pReq := &pb.ChangePasswordRequest{
		Id:       req.Id,
		Password: string(hash),
	}

	pRes, err := h.customerClient.ChangePassword(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.CustomerResponse{
		ID:           pRes.Id,
		Name:         pRes.Name,
		Address:      pRes.Address,
		License_id:   pRes.LicenseId,
		Phone_number: pRes.PhoneNumber,
		Email:        pRes.Email,
		Password:     pRes.Password,
		Active:       pRes.Active,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func NewCustomerHandler(customerClient pb.BookingCustomerClient) CustomerHandler {
	return &customerHandler{
		customerClient: customerClient,
	}
}
