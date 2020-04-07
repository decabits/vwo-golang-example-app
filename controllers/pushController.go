package controllers

import (
	"net/http"

	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/decabits/vwo-golang-sdk/api"
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
		tagKey = "hello2"
	}

	if tagValue == "" {
		tagValue = "v"
	}

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()

	result := api.Push(instance, tagKey, tagValue, userID)

	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
		"result": result,
		"tagKey": tagKey,
		"tagValue": tagValue,
	})
}
