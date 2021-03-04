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
	api.POST("/user/register", handlerService.RegisterUserHandler)
	api.POST("/user/login", handlerService.LoginUserHandler)
	api.POST("/user/check-email", handlerService.CheckEmailAvailabilityHandler)
	api.POST("/user/upload/avatars", handler.AuthMiddleware(auth, service), handlerService.UploadAvatarHandler)

	api.GET("/campaign/list", handlerService.ListCampaignHandler)
	api.GET("/campaign/detail/:uuid", handlerService.DetailCampaignHandler)
	api.POST("/campaign/create", handler.AuthMiddleware(auth, service), handlerService.CreateCampaignHandler)
	api.PUT("/campaign/update/:uuid", handler.AuthMiddleware(auth, service), handlerService.UpdateCampaignHandler)

	router.Run()
}
