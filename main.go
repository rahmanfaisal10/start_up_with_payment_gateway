package main

import (
	"fmt"
	"log"

	"bwastartup/handler"
	"bwastartup/user"

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

	fmt.Println("Connection to Database Success")
	//init repository, service and handler
	userRepository := user.InitRepository(db)
	userServices := user.InitService(userRepository)
	userHandler := handler.InitHandler(userServices)

	router := gin.Default()

	api := router.Group("api/v1")
	api.POST("/register", userHandler.RegisterUserHandler)

	router.Run()

}