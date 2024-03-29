package apis

import (
	"daily_golang/src"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SimpleTest(c *gin.Context) {
	msg := fmt.Sprintf("Call successful 200")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func SimpleQuery(c *gin.Context) {
	msg := fmt.Sprintf("Call successful 200 query ")
	user, _ := src.QueryUser(1)
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"user": user,
	})
}
