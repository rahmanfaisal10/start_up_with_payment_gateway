package main

import (
	"bwastartup/auth"
	"bwastartup/pkg/handler"
	"bwastartup/pkg/model"
	"bwastartup/pkg/repository"
	"bwastartup/pkg/service"
	"log"

	"github.com/gin-contrib/cors"
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

	//migration struct to table
	db.AutoMigrate(model.User{}, model.Campaign{}, model.CampaignImage{}, model.Transaction{})

	//init repository, service and handler
	repository := repository.InitRepository(db)
	auth := auth.InitAuthorization()
	service := service.InitService(repository)
	handlerService := handler.InitHandler(service, auth)

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Logger(), gin.Recovery())

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

	api.POST("campaign-image/create", handler.AuthMiddleware(auth, service), handlerService.CreateCampaignImageHandler)

	api.GET("transaction/campaign/:campaign_id", handler.AuthMiddleware(auth, service), handlerService.ListTransactionByCampaignIDHandler)
	api.GET("transaction/campaigns/:user_id", handler.AuthMiddleware(auth, service), handlerService.ListTransactionByUserIDHandler)
	router.Run()
}
