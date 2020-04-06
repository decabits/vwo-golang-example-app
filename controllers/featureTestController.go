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
	userID := util.GetRandomUser()
	campaignKey := config.GetString("featureTestCampaignKey")

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
