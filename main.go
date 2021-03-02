package main

import (
	"bwastartup/auth"
	"bwastartup/pkg/handler"
	"bwastartup/pkg/repository"
	"bwastartup/pkg/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	//init repository, service and handler
	repository := repository.InitRepository(db)
	auth := auth.InitAuthorization()
	service := service.InitService(repository)
	handlerService := handler.InitHandler(service, auth)

	router := gin.Default()

	router.Static("assets/images", "./assets/images")

	api := router.Group("api/v1")
	api.POST("/register", handlerService.RegisterUserHandler)
	api.POST("/login", handlerService.LoginUserHandler)
	api.POST("/check-email", handlerService.CheckEmailAvailabilityHandler)
	api.POST("/avatars", handler.AuthMiddleware(auth, service), handlerService.UploadAvatarHandler)

	api.GET("/list-campaigns", handlerService.ListCampaignHandler)
	api.GET("/detail-campaign/:uuid", handlerService.DetailCampaignHandler)

	router.Run()
}
