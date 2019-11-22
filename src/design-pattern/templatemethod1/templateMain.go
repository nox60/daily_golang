package main

import (
	"daily_golang/src"
)

func main() {
	ftpDownloader := src.FtpDownloader{}
	template := src.Template{}
	template.Implement = &ftpDownloader
	template.Download("", "")
}
