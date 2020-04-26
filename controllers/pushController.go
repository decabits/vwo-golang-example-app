package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
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

	result := instance.Push(tagKey, tagValue, userID)

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "push.html", gin.H{
		"settingsFile": string(settingsFile),
		"tagKey":       tagKey,
		"tagValue":     tagValue,
		"userID":       userID,
		"result":       result,
	})
}
