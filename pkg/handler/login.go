package handler

import (
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) LoginUserHandler(c *gin.Context) {
	req := new(request.LoginUserInput)

	err := c.ShouldBind(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("your account cannot to login", "failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	result, err := h.service.LoginUserService(*req)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		resp := response.ResponseAPI("your account cannot to login", "failed", http.StatusBadRequest, errorMessage)
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

	resp := response.ResponseAPI("success to Login", "success", http.StatusOK, token)
	c.JSON(http.StatusAccepted, resp)
}
