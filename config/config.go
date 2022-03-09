// Package config 配置类
package config

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

const (
	TYPE_IPV4       = "ipv4"
	TYPE_IPV4_LOCAL = "ipv4-local" // ipv4 lan口的内网地址
	TYPE_IPV6       = "ipv6"
	TYPE_DEFAULT    = "default" // 走默认接口 ipv4的批量修改一批sub为同一公网ip
)

const (
	BASE_URL   = "https://wx.scjtqs.com"
	BASE_URLV6 = "https://wx.mobyds.com" // 自动识别ipv6用的接口
)

var (
	// LittleC   string
	Version   string
	goversion = runtime.Version()
	ver       bool
	help      bool
)

// Config 配置信息
type Config struct {
	UserID    int64  `json:"user_id" yaml:"user_id"`       // 用户名ID
	Token     string `json:"token" yaml:"token"`           // 验证token
	Type      string `json:"type" yaml:"type"`             // ddns类型 ipv4 ipv6
	Sub       string `json:"sub" yaml:"sub"`               // 字域名
	Domain    string `json:"domain" yaml:"domain"`         // 主域名
	ExtScript string `json:"ext_script" yaml:"ext_script"` // 使用额外shell来获取ip地址
}

// Parse parse flags
func Parse() *Config {
	// wd, _ := os.Getwd()
	// dc := path.Join(wd, "config.yml")
	c := Config{}
	// flag.StringVar(&LittleC, "c", dc, "configuration filename")
	flag.Int64Var(&c.UserID, "id", 0, "你的用户ID")
	flag.StringVar(&c.Token, "token", "", "你的token秘钥")
	flag.StringVar(&c.Sub, "sub", "", "sub子域名 eg: www 仅 type 为 ipv4 和 ipv6 的时候有用")
	flag.StringVar(&c.Domain, "domain", "", "主域名 eg: scjtqs.com 仅 type 为 ipv4 和 ipv6 的时候有用")
	flag.StringVar(&c.Type, "type", "default", `使用类型：
default 默认类型 使用公网ipv4地址替换web网站上配置的所有的域名信息
ipv4  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前网络的公网IP
ipv4-local  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前设备的 内网ipv4地址 
ipv6  独立使用。更新 sub.domain 的域名的ipv6地址，默认情况下为使用当前设备对外的IPV6地址
`)
	flag.StringVar(&c.ExtScript, "extscript", "", "使用额外脚本来获取ipv4/ipv6地址。不使用留空")
	flag.BoolVar(&ver, "v", false, "show version info")
	flag.BoolVar(&help, "h", false, "show version info")
	flag.Parse()
	// 版本显示
	if ver || help {
		showVersion()
		flag.PrintDefaults()
		os.Exit(0)
	}
	return &c
}

// showVersion 显示帮助菜单
func showVersion() {
	fmt.Printf(`scjtqs ddns client by golang:
Version: %s
GO Version: %s

Usage:
./ddns -id 1 -token xxxxx
./ddns -id 1 -token xxxxx -type ipv4 -sub www -domain scjtqs.com
./ddns -id 1 -token xxxxx -type ipv6 -sub ipv6 -domain scjtqs.com -extscript /path/ipv6.sh

参数说明：
`, Version, goversion)
}
