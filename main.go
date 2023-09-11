package main

import (
	"net/http"

	"github.com/Badiru27/hngx-stage-2/configs"
	"github.com/Badiru27/hngx-stage-2/routes"
	"github.com/gin-gonic/gin"
)

 func main(){

	route := gin.Default();

	route.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to stage2 Task 💸💸💸")
	});

	configs.ConnectToDB()

	routes.UserRoute(route)

	
	route.Run()
 }