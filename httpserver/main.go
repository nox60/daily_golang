package main

import (
	. "daily_golang/httpserver/utils"
)

func main() {
	//defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8080")
}
