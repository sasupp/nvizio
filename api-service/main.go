package main

import (
	"log"
	"os"

	"xtrinio.com/app"
)

func usage() {
	log.Printf("srv <cluster ips>...")
}

func main() {

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	app.Server(os.Args[1:])
}

// Build app with
// go build -ldflags="-s -w" main.go && upx --best --lzma main
