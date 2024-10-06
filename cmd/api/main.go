package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Player struct{
	Name			string		`json:"name"	binding:"required"`
	Age				int			`json:"age"		binding:"required"`
	CreationTime	time.Time	`json:"creation_time"`
}

func main(){
	
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", func(c *gin.Context){
		var player Player
		if err := c.BindJSON(&player); err != nil{
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	})

	log.Fatal(ginEngine.Run(":8001"))
}