// Package config 配置类
package config

import "regexp"

const (
	TYPE_IPV4       = "ipv4"
	TYPE_IPV4_LOCAL = "ipv4-local" // ipv4 lan口的内网地址
	Type_IPV4_OUT   = "ipv4-out"   // 使用外部获取Ipv4地址提交
	TYPE_IPV6_LOCAL = "ipv6-local" // ipv6 lan口的内网地址
	TYPE_IPV6       = "ipv6"
	TYPE_DEFAULT    = "default" // 走默认接口 ipv4的批量修改一批sub为同一公网ip
)

const (
	BASE_URL   = "https://wx.scjtqs.com"
	BASE_URLV6 = "https://wx6.scjtqs.com" // 自动识别ipv6用的接口
)

// Config 配置信息
type Config struct {
	UserID    int64  `json:"user_id" yaml:"user_id"`       // 用户名ID
	Token     string `json:"token" yaml:"token"`           // 验证token
	Type      string `json:"type" yaml:"type"`             // ddns类型 ipv4 ipv6
	Sub       string `json:"sub" yaml:"sub"`               // 字域名
	Domain    string `json:"domain" yaml:"domain"`         // 主域名
	ExtScript string `json:"ext_script" yaml:"ext_script"` // 使用额外shell来获取ip地址
	Server    bool   `json:"server" yaml:"server"`         // 是否以server方式免crontab方式运行
}

// Ipv4Reg IPv4正则
var Ipv4Reg = regexp.MustCompile(`((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])`)

// Ipv6Reg IPv6正则
var Ipv6Reg = regexp.MustCompile(`((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))`)

var DefaultIpv4Out = "https://myip.ipip.net,https://ddns.oray.com/checkip,https://ip.3322.net,https://4.ipw.cn"
