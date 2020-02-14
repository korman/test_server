package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	fmt.Print(c.Request.Header)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}