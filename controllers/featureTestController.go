package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/gin-gonic/gin"
)

// FeatureTestController function uses the configuration values and VWO instance to check 
// whether a feature is enabled or not, gets the value of the variable for the given user and displayes the html output
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
		instance.Logger.Error(err)
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
