package apis

import (
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
