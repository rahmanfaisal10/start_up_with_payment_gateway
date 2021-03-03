package handler

import (
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterUserHandler(c *gin.Context) {
	req := new(request.RegisterUserInput)

	err := c.ShouldBindJSON(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("your account cannot to register", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	result, err := h.service.RegisterUserService(*req)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		resp := response.ResponseAPI("your account cannot to register", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	token, err := h.authorization.GenerateToken(*result)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		resp := response.ResponseAPI(" register failed", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	formatter := response.FormaterUserResponse(*result, token)
	resp := response.ResponseAPI("your account has been register", "success", http.StatusCreated, formatter)
	c.JSON(http.StatusCreated, resp)
}
