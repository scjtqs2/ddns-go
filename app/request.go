package app

import (
	"context"
	"errors"
	"github.com/guonaihong/gout"
	"github.com/scjtqs2/ddns-go/config"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Get get方式的http请求
func Get(URL string) (rsp string, err error) {
	err = gout.GET(URL).BindBody(&rsp).Do()
	return
}

// GetIPV4  IPv4 get方式的http请求
func GetIPV4(URL string) (rsp string, err error) {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	var zeroDialer net.Dialer
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return zeroDialer.DialContext(ctx, "tcp4", addr)
	}
	httpClient.Transport = transport
	err = gout.New(httpClient).GET(URL).BindBody(&rsp).Do()
	return
}

// GetOutBoundIP 获取本地IPV4
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "223.5.5.5:53")
	if err != nil {
		log.Error(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// GetPublicIPV4 获取公网IPV4
func GetPublicIPV4() (ip string, err error) {
	url := config.BASE_URL + "/ip.php"
	ip, err = GetIPV4(url)
	if ip == "unknow" {
		return "", errors.New("no ipv4")
	}
	return ip, err
}

func GetPublicIPV6() (ip string, err error) {
	url := config.BASE_URLV6 + "/ip6.php"
	ip, err = GetIPV4(url)
	if ip == "no ipv6" {
		return "", errors.New("no ipv6")
	}
	return ip, err
}

// GetOutBoundIPV6 获取本地IPV6
func GetOutBoundIPV6() (ip string, err error) {
	conn, err := net.Dial("udp", "[2400:3200::1]:53")
	if err != nil {
		// fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	iparr := regexp.MustCompile(`^\[([\w+:]+)\]:\d+`).FindAllStringSubmatch(localAddr.String(), -1)
	if len(iparr) < 1 {
		err = errors.New("获取ipv6地址失败")
	} else {
		ip = iparr[0][1]
	}
	return
}

// // Test 打印所有的本地addr
// func Test() {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, address := range addrs {
// 		// 检查ip地址判断是否回环地址
// 		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
// 			if ipnet.IP.To4() != nil {
// 				// ipv4地址打印
// 				fmt.Println(ipnet.IP.String())
// 			}
// 			if ipnet.IP.To16() != nil {
// 				// ipv6地址打印
// 				fmt.Println(ipnet.IP.String())
// 			}
// 		}
// 	}
// }
