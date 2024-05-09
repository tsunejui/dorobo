package main

import (
	"dorobo/web"
	"log"
	"os"
)

func main() {
	svc, err := web.New()
	if err != nil {
		log.Fatalf("failed to run web server: %v", err)
	}
	svc.Run() // run web server
	os.Exit(0)
}
