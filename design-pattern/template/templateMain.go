package main

import (
	. "./template"
)

func main() {
	ftpDownloader := FtpDownloader{}
	template := Template{}
	template.Implement = &ftpDownloader
}
