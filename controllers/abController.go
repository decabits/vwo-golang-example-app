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
	userID := c.Query("userId")
	if userID == "" {
		userID = util.GetRandomUser()
	}
	// userID = "Gimmy"
	campaignKey := c.Query("cKey")
	if campaignKey == "" {
		campaignKey = config.GetString("abCampaignKey")
	}
	// campaignKey = "phpab3"
	abCampaigngoalIdentifier := config.GetString("abCampaignGoalIdentifier")
	// abCampaigngoalIdentifier = "custom"

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

	var settings schema.SettingsFile
	settings = instance.SettingsFile

	c.HTML(http.StatusOK, "ab.html", gin.H{
		"campaign_type":               "Visual_AB",
		"settings_file":               settings,
		"ab_campaign_key":             campaignKey,
		"ab_campaign_goal_identifier": "",
		"custom_variables":            "",
		"user_id":                     userID,
		"is_part_of_campaign":         isPartOfCampaign,
		"variation_name":              variationName,
		"button":                      `<button class="v1">Control</button>`,
		"track":                       track,
	})
}
