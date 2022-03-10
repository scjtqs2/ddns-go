package main

import (
	"github.com/kardianos/service"
	"github.com/scjtqs2/ddns-go/app"
	"log"
	"os"
	"os/signal"
	"time"
)

// program 后台服务管理
type program struct {
	Client *app.Client
}

// Start 脚本启动
func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

// run 后台运行的启动
func (p *program) run() {

	err := p.Client.Run()
	if err != nil {
		log.Fatalln(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signle := <-c
	log.Printf("quit,Got signal:", signle)
	os.Exit(1)

}

// Stop service stop
func (p *program) Stop(s service.Service) error {
	if p.Client != nil {
		p.Client.Cron.Stop()
	}
	log.Printf("stoping")
	<-time.After(time.Second * 2)
	return nil
}
