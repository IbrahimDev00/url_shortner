package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	setUpRouter(router)

	setUpRouter(router)

	port := os.Getenv("APP_PORT")

	if port == ""{
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}

func setUpRouter(router *gin.Engine){
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.Del("/api/v1/:shortID", routes.DeleteURL")
	router.PUT("/api/v1/:shortID", routes.EditURL")		
	router.POST('api/v1/addTag', routes.AddTag')
}