package handler

import (
	"bwastartup/response"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func InitHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	req := new(user.RegisterUserInput)

	err := c.ShouldBind(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("your account cannot to register", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	result, err := h.service.RegisterUserService(*req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("your account cannot to register", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	formatter := response.FormaterUserResponse(*result)
	resp := response.ResponseAPI("your account has been register", "success", http.StatusCreated, formatter)
	c.JSON(http.StatusCreated, resp)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	req := new(user.LoginUserInput)

	err := c.ShouldBind(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("your account cannot to login", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	_, err = h.service.LoginUserService(*req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("your account cannot to login", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.ResponseAPI("success to Login", "success", http.StatusAccepted, nil)
	c.JSON(http.StatusAccepted, resp)
}

func (h *userHandler) CheckEmailAvailabilityHandler(c *gin.Context) {
	req := new(user.CheckEmailAvailable)
	err := c.ShouldBind(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("email checking failed", "error ", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	isEmailAvailable, err := h.service.CheckEmailAvailabilityService(*req)
	if err != nil {
		errorMessage := gin.H{"error": "server error"}
		resp := response.ResponseAPI("email checking failed", "error ", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	resp := response.ResponseAPI(metaMessage, "success", http.StatusAccepted, isEmailAvailable)
	c.JSON(http.StatusAccepted, resp)
}
