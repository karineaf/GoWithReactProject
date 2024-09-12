package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/karineaf/GoWithReactProject/src/controller/routes" 
)

func main(){

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal((err))
	}
}