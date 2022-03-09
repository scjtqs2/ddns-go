# 说明

https://wx.scjtqs.com/ddns 的客户端 by golang

```shell
scjtqs ddns client by golang:
Version: 
GO Version: go1.17.6

Usage:
./ddns -id 1 -token xxxxx
./ddns -id 1 -token xxxxx -type ipv4 -sub www -domain scjtqs.com
./ddns -id 1 -token xxxxx -type ipv6 -sub ipv6 -domain scjtqs.com -extscript /path/ipv6.sh

参数说明：
  -domain string
        主域名 eg: scjtqs.com 仅 type 为 ipv4 和 ipv6 的时候有用
  -extscript string
        使用额外脚本来获取ipv4/ipv6地址。不使用留空
  -h    show version info
  -id int
        你的用户ID
  -sub string
        sub子域名 eg: www 仅 type 为 ipv4 和 ipv6 的时候有用
  -token string
        你的token秘钥
  -type string
        使用类型：
        default 默认类型 使用公网ipv4地址替换web网站上配置的所有的域名信息
        ipv4  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前网络的公网IP
        ipv4-local  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前设备的 内网ipv4地址 
        ipv6  独立使用。更新 sub.domain 的域名的ipv6地址，默认情况下为使用当前设备对外的IPV6地址
         (default "default")
  -v    show version info

```