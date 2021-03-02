package handler

import (
	"bwastartup/auth"
	"bwastartup/pkg/response"
	"bwastartup/pkg/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(auth auth.Auth, service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")
		fmt.Println(authHeader)

		if !strings.Contains(authHeader, "Bearer") {
			response := response.ResponseAPI("Unathorization", "error", http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""

		arrayString := strings.Split(authHeader, " ")
		if len(arrayString) == 2 {
			tokenString = arrayString[1]
		}

		tokenJWT, err := auth.ValidationToken(tokenString)
		if err != nil {
			response := response.ResponseAPI("Unathorization", "error", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, ok := tokenJWT.Claims.(jwt.MapClaims)
		if !ok || !tokenJWT.Valid {
			response := response.ResponseAPI("Unathorization", "error", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claims["user_id"].(string)

		user, err := service.GetUserByIDService(userID)
		if err != nil {
			response := response.ResponseAPI("Unathorization", "error", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("current_user", user)
	}
}
