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

// ABController ...
func ABController(c *gin.Context) {
	config := config.GetConfig()
	userID := util.GetRandomUser()
	// userID = "Faizan"
	campaignKey := config.GetString("abCampaignKey")
	abCampaigngoalIdentifier := config.GetString("abCampaignGoalIdentifier")

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()
	options := schema.Options{}

	isPartOfCampaign := false
	variationName := api.ActivateWithOptions(instance, campaignKey, userID, options)
	if variationName != "" {
		isPartOfCampaign = true
	}

	track := api.TrackWithOptions(instance, campaignKey, userID, abCampaigngoalIdentifier, options)

	c.JSON(http.StatusOK, gin.H{
		"userID":           userID,
		"variation":        variationName,
		"campaign":         campaignKey,
		"isPartOfCampaign": isPartOfCampaign,
		"track":            track,
	})
}
