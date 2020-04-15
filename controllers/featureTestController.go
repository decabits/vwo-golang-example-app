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

	// c.JSON(http.StatusOK, gin.H{
	// 	"userID":      userID,
	// 	"campaignKey": campaignKey,
	// 	"isEnabled":   isEnabled,
	// 	"string": gin.H{"key": stringVariable,
	// 		"value": strValue},
	// 	"integer": gin.H{"key": intVariable,
	// 		"value": intValue},
	// 	"boolean": gin.H{"key": boolVariable,
	// 		"value": boolValue},
	// 	"double": gin.H{"key": doubleVariable,
	// 		"value": dubValue},
	// })

	var settings schema.SettingsFile
	settings = instance.SettingsFile

	c.HTML(http.StatusOK, "feature-test.html", gin.H{
		"campaign_type": "Feature_Test",
		"settings_file": settings,
		"feature_campaign_key": campaignKey,
		"feature_campaign_goal_identifier": "",
		"custom_variables": "",
		"user_id": userID,
		"is_user_part_of_feature_campaign": isEnabled,
		"boolean_variable": boolValue,
		"integer_variable": intValue,
		"double_variable": dubValue,
		"string_variable": strValue,
	})
}
