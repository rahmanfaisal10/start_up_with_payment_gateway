package handler

import (
	"bwastartup/pkg/model"
	"bwastartup/pkg/request"
	"bwastartup/pkg/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) RegisterUserHandler(c *gin.Context) {
	req := new(request.RegisterUserInput)

	err := c.ShouldBind(req)
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

	token, err := h.authorization.GenerateToken(int(result.ID))
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

	token, err := h.authorization.GenerateToken(int(result.ID))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		resp := response.ResponseAPI(" register failed", "failed", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := response.ResponseAPI("success to Login", "success", http.StatusAccepted, token)
	c.JSON(http.StatusAccepted, resp)
}

func (h *handler) CheckEmailAvailabilityHandler(c *gin.Context) {
	req := new(request.CheckEmailAvailable)
	err := c.ShouldBind(req)
	if err != nil {
		errorMessage := gin.H{"error": response.ErrorValidationResponse(err)}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	isEmailAvailable, err := h.service.CheckEmailAvailabilityService(*req)
	if err != nil {
		errorMessage := gin.H{"error": "server error"}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusUnprocessableEntity, errorMessage)
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

func (h *handler) UploadAvatarHandler(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	path := fmt.Sprintf("assets/images/%s", file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	cuurentUser := c.MustGet("current_user").(model.User)
	userID := int(cuurentUser.ID)
	fmt.Println("berapa nilai userID", userID)

	_, err = h.service.SaveAvatarService(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		resp := response.ResponseAPI("email checking failed", "error", http.StatusBadRequest, data)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data := gin.H{"is_uploaded": true}
	resp := response.ResponseAPI("email checking failed", "success", http.StatusOK, data)
	c.JSON(http.StatusOK, resp)
}
