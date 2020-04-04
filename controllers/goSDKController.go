package controllers

import (
	"net/http"

	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/gin-gonic/gin"
)

// GoSDKController ...
func GoSDKController(c *gin.Context) {
	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()

	c.JSON(http.StatusOK, instance.SettingsFile)
}
