package utils

import (
	. "daily_golang/httpserver/apis"
	"database/sql"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func InitRouter() *gin.Engine {
	//初始化连接池
	/*
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
		if err != nil{
			log.Fatalln(err)
		}
		defer db.Close()

		//很多参数使用环境变量
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(5)

		if err := db.Ping(); err != nil{
			log.Fatalln(err)
		}
	*/
	router := gin.Default()
	router.GET("/simpleTest", SimpleTest)
	router.GET("/simpleQuery", SimpleQuery)

	return router
}
