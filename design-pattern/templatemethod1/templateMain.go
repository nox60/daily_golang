package main

import (
	"./template"
	"fmt"
)

func main() {
	ftpDownloader := template.FtpDownloader{}
	template := template.Template{}
	template.Implement = &ftpDownloader
	template.Download("", "")

	httpDownloader := template.HttpDownloader{}
	fmt.Println(httpDownloader)
}
