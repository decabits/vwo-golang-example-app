package controllers

import (
	"net/http"

	"github.com/Piyushhbhutoria/vwo-go-sdk/api"
	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/gin-gonic/gin"
)

func ABController(c *gin.Context) {
	config := config.GetConfig()
	userID := util.GetRandomUser()

	vwo := models.VWO{}
	instance := vwo.GetVWOInstance()
	variationName := api.Activate(instance, config.GetString("abCampaignKey"), userID)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
