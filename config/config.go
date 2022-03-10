// Package config 配置类
package config

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
