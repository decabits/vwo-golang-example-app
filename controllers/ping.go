package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusController struct{}

func (s StatusController) Status(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}
