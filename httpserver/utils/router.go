package utils

import (
	. "daily_golang/httpserver/apis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/simpleTest", SimpleTest)
	return router
}
