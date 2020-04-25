package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
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
	campaignKey := c.Query("cKey")
	if campaignKey == "" {
		campaignKey = config.GetString("abCampaignKey")
	}
	abCampaigngoalIdentifier := config.GetString("abCampaignGoalIdentifier")
	// abCampaigngoalIdentifier = "custom"

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()
	options := schema.Options{}

	isPartOfCampaign := false
	variationName := instance.ActivateWithOptions(campaignKey, userID, options)
	if variationName != "" {
		isPartOfCampaign = true
	}

	track := instance.TrackWithOptions(campaignKey, userID, abCampaigngoalIdentifier, options)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	var class string
	if variationName == "Control" {
		class = "v1"
	} else if variationName == "Variation-1" {
		class = "v2"
	} else {
		class = "v3"
	}

	c.HTML(http.StatusOK, "ab.html", gin.H{
		"campaignType":             "VISUAL_AB",
		"settingsFile":             string(settingsFile),
		"abCampaignKey":            campaignKey,
		"abCampaignGoalIdentifier": abCampaigngoalIdentifier,
		"customVariables":          options.CustomVariables,
		"userID":                   userID,
		"isPartOfCampaign":         isPartOfCampaign,
		"variationName":            variationName,
		"buttonClass":              class,
		"track":                    track,
	})
}
