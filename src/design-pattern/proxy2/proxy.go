package main

import (
	"daily_golang/src"
)

func main() {
	//fileAccessorImpl := fileAccessorImpl{} 用户不能访问到 fileAccessorImpl, 利用包权限隔离，只能通过下面的代理实现访问
	userFileAccessProxy := src.FileAccessorProxy{UserType: "user"}
	guestFileAccessProxy := src.FileAccessorProxy{UserType: "guest"}
	userFileAccessProxy.ReadFile("/opt/user/file")
	guestFileAccessProxy.ReadFile("/opt/user/file")
}
