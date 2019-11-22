package main

import "daily_golang/design-pattern/templatemethod1/template"

func main() {
	ftpDownloader := template.FtpDownloader{}
	template := template.Template{}
	template.Implement = &ftpDownloader
	template.Download("", "")
}
