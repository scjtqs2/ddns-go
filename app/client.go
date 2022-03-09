// Package app ddns的应用类
package app

import (
	"fmt"
	"github.com/scjtqs2/ddns-go/config"
)

// Client 客户端的实体类
type Client struct {
	Config *config.Config
}

// NewClient 新建客户端
func NewClient(cof *config.Config) *Client {
	return &Client{
		Config: cof,
	}
}

// Run 运行
func (c *Client) Run() {
	switch c.Config.Type {
	case config.TYPE_IPV4:
		c.ipv4Run()
	case config.TYPE_IPV6:
		c.ipv6Run()
	case config.TYPE_DEFAULT:
		c.defaultRun()
	}
}

// defaultRun 默认方式运行
func (c *Client) defaultRun() {
	url := fmt.Sprintf("%s/ddns/client/id/%d/token/%s", config.BASE_URL, c.Config.UserID, c.Config.Token)
	rsp, err := Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("default rsp:%s", rsp)
}

// ipv4Run ipv4模式运行
func (c *Client) ipv4Run() {
	var (
		ipv4 string
		err  error
	)
	switch c.Config.ExtScript {
	case "":
		ipv4, err = GetPublicIPV4()
	default:
		ipv4, err = Command(c.Config.ExtScript)
	}
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%s/ddns/client/ipv4?id=%d&token=%s&sub=%s&domain=%s&ipv4=%s", config.BASE_URL, c.Config.UserID, c.Config.Token, c.Config.Sub, c.Config.Domain, ipv4)
	rsp, err := Get(url)
	fmt.Printf("ipv4 rsp:%s", rsp)
}

// ipv6Run ipv6模式运行
func (c *Client) ipv6Run() {
	var (
		ipv6 string
		err  error
	)
	switch c.Config.ExtScript {
	case "":
		ipv6, err = GetOutBoundIPV6()
	default:
		ipv6, err = Command(c.Config.ExtScript)
	}
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%s/ddns/client/ipv6?id=%d&token=%s&sub=%s&domain=%s&ipv6=%s", config.BASE_URL, c.Config.UserID, c.Config.Token, c.Config.Sub, c.Config.Domain, ipv6)
	rsp, err := Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ipv6 rsp:%s", rsp)
}
