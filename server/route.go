package server

import (
	"github.com/decabits/vwo-golang-example-app/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	ping := new(controllers.StatusController)
	router.GET("/ping", ping.Status)
	router.GET("/ab", controllers.ABController)

	return router
}
