package controllers

import (
	"net/http"

	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/decabits/vwo-golang-sdk/api"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/gin-gonic/gin"
)

func ABController(c *gin.Context) {
	config := config.GetConfig()
	userID := util.GetRandomUser()
	campaignKey := config.GetString("abCampaignKey")

	vwo := models.VWO{}
	instance := vwo.GetVWOInstance()
	options := schema.Options{}

	variationName := api.ActivateWithOptions(instance,campaignKey, userID, options)
	// variationName := api.Activate(instance,campaignKey, userID)

	c.JSON(http.StatusOK, gin.H{
		"userID":    userID,
		"variation": variationName,
	})
}
