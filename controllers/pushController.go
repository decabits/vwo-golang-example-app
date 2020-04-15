package controllers

import (
	"net/http"

	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/decabits/vwo-golang-sdk/api"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/gin-gonic/gin"
)

// PushController ...
func PushController(c *gin.Context) {
	userID := c.Query("userId")
	tagKey := c.Query("tagKey")
	tagValue := c.Query("tagValue")
	if userID == "" {
		userID = util.GetRandomUser()
	}

	if tagKey == "" {
		tagKey = "tempKey"
	}

	if tagValue == "" {
		tagValue = "tempVal"
	}

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()

	result := api.Push(instance, tagKey, tagValue, userID)

	var settings schema.SettingsFile
	settings = instance.SettingsFile

	c.HTML(http.StatusOK, "push.html", gin.H{
		"settings_file": settings,
		"tag_key":       tagKey,
		"tag_value":     tagValue,
		"user_id":       userID,
		"result":        result,
	})
}
