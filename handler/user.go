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
		errorMessage := response.ErrorValidationResponse(err)
		resp := response.ResponseAPI("your account cannot to register", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	result, err := h.service.RegisterUser(*req)
	if err != nil {
		resp := response.ResponseAPI("your account cannot to register", "failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	formatter := response.FormaterUserResponse(*result)
	resp := response.ResponseAPI("your account has been register", "success", http.StatusCreated, formatter)
	c.JSON(http.StatusCreated, resp)
}
