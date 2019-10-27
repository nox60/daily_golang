package dbaccess

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

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

func newPool() *sql.DB {
	log.Println("Init database...")
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "xxxx.COM"
	cfg.Net = "tcp"
	cfg.Addr = "www.xxxx.cn:18992"
	cfg.DBName = "cjfh"
	dsn := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

//连接池在一开始初始化，后面不会再实例化
var pool = newPool()
