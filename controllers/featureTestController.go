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

// FeatureTestController ...
func FeatureTestController(c *gin.Context) {
	config := config.GetConfig()
	userID := c.Query("userId")
	if userID == "" {
		userID = util.GetRandomUser()
	}
	// userID = "Faizan"
	campaignKey := c.Query("cKey")
	if campaignKey == "" {
		campaignKey = config.GetString("featureTestCampaignKey")
	}

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()
	options := schema.Options{}

	var (
		stringVariable = "string2"
		intVariable    = "int2"
		boolVariable   = "bool2"
		doubleVariable = "float2"
	)

	isEnabled := api.IsFeatureEnabled(instance, campaignKey, userID, options)

	strValue := api.GetFeatureVariableValue(instance, campaignKey, stringVariable, userID, options)
	intValue := api.GetFeatureVariableValue(instance, campaignKey, intVariable, userID, options)
	boolValue := api.GetFeatureVariableValue(instance, campaignKey, boolVariable, userID, options)
	dubValue := api.GetFeatureVariableValue(instance, campaignKey, doubleVariable, userID, options)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		instance.Logger.Error(err)
	}

	c.HTML(http.StatusOK, "featureTest.html", gin.H{
		"campaignType":                  "FEATURE_TEST",
		"settingsFile":                  string(settingsFile),
		"featureCampaignKey":            campaignKey,
		"featureCampaignGoalIdentifier": config.GetString("abCampaignGoalIdentifier"),
		"customVariables":               options.CustomVariables,
		"userID":                        userID,
		"isUserPartOfFeatureCampaign":  isEnabled,
		"booleanVariable":               boolValue,
		"integerVariable":               intValue,
		"doubleVariable":                dubValue,
		"stringVariable":                strValue,
	})
}
