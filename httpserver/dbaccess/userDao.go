package dbaccess

import (
	"daily_golang/httpserver/models"
	"database/sql"
	"log"
)

func QueryUser(id int64) (models.User, error) {
	row := pool.QueryRow("select `accountId`, `userName` from `tb_users` where `accountId` = ?", id)

	user := models.User{}
	if err := row.Scan(&user.ID, &user.UserName); err != nil {
		if err == sql.ErrNoRows {
			return user, nil // 返回 (*User)(nil)
		}
		return user, err
	}
	log.Println("result")
	log.Println(user)
	return user, nil
}
