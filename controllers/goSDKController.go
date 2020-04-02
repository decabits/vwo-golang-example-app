package controllers

import (
	"net/http"

	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/gin-gonic/gin"
)

func GoSDKController(c *gin.Context) {
	vwo := models.VWO{}
	instance := vwo.GetVWOInstance()

	c.JSON(http.StatusOK, instance.SettingsFile)
}
