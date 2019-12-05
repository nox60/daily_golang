package main

import (
	. "./templateMethod"
)

func main() {
	ftpDownloader := FtpDownloader{}
	template := Template{}
	template.Implement = &ftpDownloader
	template.Download("", "")
}
