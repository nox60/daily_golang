package main

import (
	"daily_golang/src"
)

func main() {
	//defer db.SqlDB.Close()
	router := src.InitRouter()
	router.Run(":8080")
}
