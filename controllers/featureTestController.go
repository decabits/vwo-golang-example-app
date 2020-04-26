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

	var (
		stringVariable = "string2"
		intVariable    = "int2"
		boolVariable   = "bool2"
		doubleVariable = "float2"
	)

	isEnabled := instance.IsFeatureEnabled(campaignKey, userID, nil)

	strValue := instance.GetFeatureVariableValue(campaignKey, stringVariable, userID, nil)
	intValue := instance.GetFeatureVariableValue(campaignKey, intVariable, userID, nil)
	boolValue := instance.GetFeatureVariableValue(campaignKey, boolVariable, userID, nil)
	dubValue := instance.GetFeatureVariableValue(campaignKey, doubleVariable, userID, nil)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "featureTest.html", gin.H{
		"campaignType":                  "FEATURE_TEST",
		"settingsFile":                  string(settingsFile),
		"featureCampaignKey":            campaignKey,
		"featureCampaignGoalIdentifier": config.GetString("abCampaignGoalIdentifier"),
		"customVariables":               nil,
		"userID":                        userID,
		"isUserPartOfFeatureCampaign":   isEnabled,
		"booleanVariable":               boolValue,
		"integerVariable":               intValue,
		"doubleVariable":                dubValue,
		"stringVariable":                strValue,
	})
}
