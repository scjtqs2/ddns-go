package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/kardianos/service"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/scjtqs2/ddns-go/app"
	"github.com/scjtqs2/ddns-go/config"
	"github.com/scjtqs2/utils/util"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	// LittleC   string
	Version   string
	goversion = runtime.Version()
	ver       bool
	help      bool
	install   bool
	uninstall bool
	cfg       *config.Config
	action    string
)

func init() {
	logFormatter := &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%time%] [%lvl%]: %msg% \n",
	}
	cfg = Parse()
	current, err := GetCurrentPath()
	if err != nil {
		panic(err)
	}
	w, err := rotatelogs.New(path.Join(current, "logs", "%Y-%m-%d.log"), rotatelogs.WithRotationTime(time.Hour*24))
	if err != nil {
		log.Errorf("rotatelogs init err: %v", err)
		panic(err)
	}
	LogLevel := "info"
	log.AddHook(util.NewLocalHook(w, logFormatter, util.GetLogLevel(LogLevel)...))
}

func main() {
	svcConfig := &service.Config{
		Name:        "ddns-go",            // 服务显示名称
		DisplayName: "ddns-go",            // 服务名称
		Description: "scjtqs 的 ddns服务客户端", // 服务描述
		Arguments: []string{
			"-id",
			strconv.FormatInt(cfg.UserID, 10),
			"-token",
			cfg.Token,
			"-type",
			cfg.Type,
		},
	}
	if install || uninstall {
		cfg.Server = true
	}
	if cfg.Server {
		svcConfig.Arguments = append(svcConfig.Arguments, "-server")
	}
	if cfg.Sub != "" {
		svcConfig.Arguments = append(svcConfig.Arguments, "-sub", cfg.Sub)
	}
	if cfg.Domain != "" {
		svcConfig.Arguments = append(svcConfig.Arguments, "-domain", cfg.Domain)
	}
	if cfg.ExtScript != "" {
		svcConfig.Arguments = append(svcConfig.Arguments, "-extscript", cfg.ExtScript)
	}
	cli := app.NewClient(cfg)

	prg := &program{Client: cli}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if install {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("服务安装成功 \r\n")
		return
	}
	if uninstall {
		_ = s.Stop()
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("服务卸载成功 \r\n")
		return
	}
	switch action {
	case "start":
		err = service.Control(s, action)
	case "stop":
		err = service.Control(s, action)
	case "restart":
		err = service.Control(s, action)
	}
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to %s %v: %v", action, s, err))
	}
	if action != "" {
		return
	}
	if cli.Config.Server {
		err = s.Run()
		if err != nil {
			log.Errorf("ERROR: %s \r\n", err.Error())
		}
	} else {
		err = cli.Run()
		if err != nil {
			log.Errorf("ERROR: %s \r\n", err.Error())
		}
	}
}

// Parse parse flags
func Parse() *config.Config {
	// wd, _ := os.Getwd()
	// dc := path.Join(wd, "config.yml")
	c := config.Config{}
	// flag.StringVar(&LittleC, "c", dc, "configuration filename")
	flag.Int64Var(&c.UserID, "id", 0, "你的用户ID")
	flag.StringVar(&c.Token, "token", "", "你的token秘钥")
	flag.StringVar(&c.Sub, "sub", "", "sub子域名 eg: www 仅 type 为 ipv4/ipv4-local 和 ipv6 的时候有用")
	flag.StringVar(&c.Domain, "domain", "", "主域名 eg: scjtqs.com 仅 type 为 ipv4/ipv4-local 和 ipv6 的时候有用")
	flag.StringVar(&c.Type, "type", "default", `使用类型：
default 默认类型 使用公网ipv4地址替换web网站上配置的所有的域名信息
ipv4  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前网络的公网IP
ipv4-local  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前设备的 内网ipv4地址 
ipv6  独立使用。更新 sub.domain 的域名的ipv6地址，默认情况下为使用当前设备对外的IPV6地址
ipv6-local  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前设备的 内网ipv6地址 
`)
	flag.StringVar(&c.ExtScript, "extscript", "", "使用额外脚本来获取ipv4/ipv6地址。不使用留空")
	flag.BoolVar(&help, "h", false, "show this info")
	flag.BoolVar(&c.Server, "server", false, "是否以server方式免crontab方式运行")
	flag.BoolVar(&install, "install", false, "安装为服务运行 安装后就能用各自系统的服务管理工具进行管理 例如linux的systemctl")
	flag.BoolVar(&uninstall, "uninstall", false, "卸载服务")
	flag.Parse()
	action = flag.Arg(0)
	// log.Printf("action:%s", action)
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

// GetCurrentPath 获取当前文件路径
func GetCurrentPath() (string, error) {
	_, file, _, _ := runtime.Caller(1)
	abs, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(abs, "/")
	if i < 0 {
		i = strings.LastIndex(abs, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(abs[0 : i+1]), nil
}
