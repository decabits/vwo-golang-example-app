package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/decabits/vwo-golang-sdk/api"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/gin-gonic/gin"
)

// FeatureRolloutController ...
func FeatureRolloutController(c *gin.Context) {
	config := config.GetConfig()
	userID := c.Query("userId")
	if userID == "" {
		userID = util.GetRandomUser()
	}
	campaignKey := c.Query("cKey")
	if campaignKey == "" {
		campaignKey = config.GetString("featureRolloutCampaignKey")
	}

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()
	options := schema.Options{}

	isEnabled := api.IsFeatureEnabled(instance, campaignKey, userID, options)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		instance.Logger.Error(err)
	}

	c.HTML(http.StatusOK, "featureRollout.html", gin.H{
		"campaignType":                        "FEATURE_ROLLOUT",
		"settingsFile":                        string(settingsFile),
		"campaifeatureRolloutCampaignKeyType": campaignKey,
		"customVariables":                     options.CustomVariables,
		"userID":                              userID,
		"isUserPartOfFeatureRolloutCampaign":  isEnabled,
	})

}
