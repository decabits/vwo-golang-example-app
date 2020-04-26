package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
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

	isEnabled := instance.IsFeatureEnabled(campaignKey, userID, nil)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "featureRollout.html", gin.H{
		"campaignType":                        "FEATURE_ROLLOUT",
		"settingsFile":                        string(settingsFile),
		"campaifeatureRolloutCampaignKeyType": campaignKey,
		"customVariables":                     nil,
		"userID":                              userID,
		"isUserPartOfFeatureRolloutCampaign":  isEnabled,
	})

}
