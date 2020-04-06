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

// FeatureRolloutController ...
func FeatureRolloutController(c *gin.Context) {
	config := config.GetConfig()
	userID := c.Query("userId")
	if userID == "" {
		userID = util.GetRandomUser()
	}
	// userID = "Faizan"
	campaignKey := c.Query("cKey")
	if campaignKey == "" {
		campaignKey = config.GetString("featureRolloutCampaignKey")
	}

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()
	options := schema.Options{}

	var (
		stringVariable = "string1"
		intVariable    = "int1"
		boolVariable   = "bool1"
		doubleVariable = "float1"
	)

	isEnabled := api.IsFeatureEnabled(instance, campaignKey, userID, options)
	
	strValue := api.GetFeatureVariableValue(instance, campaignKey, stringVariable, userID, options)
	intValue := api.GetFeatureVariableValue(instance, campaignKey, intVariable, userID, options)
	boolValue := api.GetFeatureVariableValue(instance, campaignKey, boolVariable, userID, options)
	dubValue := api.GetFeatureVariableValue(instance, campaignKey, doubleVariable, userID, options)

	c.JSON(http.StatusOK, gin.H{
		"userID":      userID,
		"campaignKey": campaignKey,
		"isEnabled":   isEnabled,
		"string": gin.H{"key": stringVariable,
			"value": strValue},
		"integer": gin.H{"key": intVariable,
			"value": intValue},
		"boolean": gin.H{"key": boolVariable,
			"value": boolValue},
		"double": gin.H{"key": doubleVariable,
			"value": dubValue},
	})
}
