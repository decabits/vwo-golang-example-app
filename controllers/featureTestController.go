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

	isEnabled := instance.IsFeatureEnabledWithOptions(campaignKey, userID, options)

	strValue := instance.GetFeatureVariableValueWithOptions(campaignKey, stringVariable, userID, options)
	intValue := instance.GetFeatureVariableValueWithOptions(campaignKey, intVariable, userID, options)
	boolValue := instance.GetFeatureVariableValueWithOptions(campaignKey, boolVariable, userID, options)
	dubValue := instance.GetFeatureVariableValueWithOptions(campaignKey, doubleVariable, userID, options)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "featureTest.html", gin.H{
		"campaignType":                  "FEATURE_TEST",
		"settingsFile":                  string(settingsFile),
		"featureCampaignKey":            campaignKey,
		"featureCampaignGoalIdentifier": config.GetString("abCampaignGoalIdentifier"),
		"customVariables":               options.CustomVariables,
		"userID":                        userID,
		"isUserPartOfFeatureCampaign":   isEnabled,
		"booleanVariable":               boolValue,
		"integerVariable":               intValue,
		"doubleVariable":                dubValue,
		"stringVariable":                strValue,
	})
}
