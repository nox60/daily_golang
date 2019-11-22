package main

import (
	"./template"
)

func main() {
	ftpDownloader := template.FtpDownloader{}
	template := template.Template{}
	template.Implement = &ftpDownloader
	template.Download("", "")


}
