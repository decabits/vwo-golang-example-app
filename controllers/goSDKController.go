package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/gin-gonic/gin"
)

// GoSDKController function displayes the settings file
func GoSDKController(c *gin.Context) {
	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()

	settingsFile, err := json.Marshal(instance.SettingsFile)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "goSDK.html", gin.H{
		"settingsFile": string(settingsFile),
	})
}
