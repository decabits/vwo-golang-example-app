package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HomePage Controller ...
func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "homePage.html", gin.H{
		// "title": "Main website",
	})
}
