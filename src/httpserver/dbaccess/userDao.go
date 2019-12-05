package dbaccess

import (
	"daily_golang/src"
	"database/sql"
	"log"
)

func QueryUser(id int64) (src.User, error) {
	row := src.pool.QueryRow("select `account_id`, `user_name` from `t_user` where `account_id` = ?", id)

	user := src.User{}
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
