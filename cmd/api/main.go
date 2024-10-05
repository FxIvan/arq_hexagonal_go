package main

import (
	"log"

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

	
}