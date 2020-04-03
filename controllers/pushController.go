package controllers

import (
	"net/http"

	"github.com/decabits/vwo-golang-example-app/models"
	"github.com/decabits/vwo-golang-example-app/util"
	"github.com/decabits/vwo-golang-sdk/api"
	"github.com/gin-gonic/gin"
)

func PushController(c *gin.Context) {
	userID := util.GetRandomUser()

	tagKey := "hello2"
  	tagValue := "b"

	vwo := models.VWO{}
	vwo.Init()
	instance := vwo.GetVWOInstance()

	result := api.Push(instance, tagKey, tagValue, userID)

	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
		"result": result,
	})
}
