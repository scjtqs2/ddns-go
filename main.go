package main

import (
	"github.com/scjtqs2/ddns-go/app"
	"github.com/scjtqs2/ddns-go/config"
)

func main() {
	cfg := config.Parse()
	cli := app.NewClient(cfg)
	cli.Run()
}
