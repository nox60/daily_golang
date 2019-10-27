package main

import (
	. "daily_golang/httpserver/utils"
	"fmt"
)

func main() {
	//defer db.SqlDB.Close()
	router := InitRouter()
	router.Run(":8000")
	fmt.Println("test resored111222")
}
