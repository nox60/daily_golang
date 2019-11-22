package proxy

import (
	"daily_golang/src"
	"fmt"
)

type FileAccessorProxy struct {
	UserType string
}

func (f *FileAccessorProxy) ReadFile(filePath string) {
	if f.UserType == "user" {
		fmt.Printf("Current user: %s \n", f.UserType)
		fileAccessor := src.fileAccessorImpl{}
		fileAccessor.ReadFile(filePath)
	} else if f.UserType == "guest" {
		fmt.Printf("Current user: %s \n", f.UserType)
		fmt.Println("Sorry, you cannot access this file")
	}
}
