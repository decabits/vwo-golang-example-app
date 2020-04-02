package controllers

import (
	"net/http"

	"github.com/Piyushhbhutoria/vwo-go-sdk/api"
	"github.com/Piyushhbhutoria/vwo-go-sdk/schema"
	"github.com/decabits/vwo-golang-example-app/config"
	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/gin-gonic/gin"
)

func FeatureRolloutController(c *gin.Context) {
	config := config.GetConfig()
	userID := util.GetRandomUser()
	campaignKey := config.GetString("featureRolloutCampaignKey")
	
	vwo := models.VWO{}
	instance := vwo.GetVWOInstance()
	options := schema.Options{}

	var (
		stringVariable string
		intVariable string
		boolVariable string
		doubleVariable string
	)
	strValue := api.GetFeatureVariableValue(instance, campaignKey, stringVariable, userID, options)
	intValue := api.GetFeatureVariableValue(instance, campaignKey, intVariable, userID, options)
	boolValue := api.GetFeatureVariableValue(instance, campaignKey, boolVariable, userID, options)
	dubValue := api.GetFeatureVariableValue(instance, campaignKey, doubleVariable, userID, options)

	c.JSON(http.StatusOK, gin.H{
      "stringkey": stringVariable,
      "stringvalue": strValue,
      "intkey": intVariable,
      "intvalue": intValue,
      "boolkey": boolVariable,
      "boolvalue": boolValue,
      "doublekey": doubleVariable,
      "doublevalue": dubValue,
    })
}
