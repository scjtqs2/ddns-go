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
        主域名 eg: scjtqs.com 仅 type 为 ipv4/ipv4-local 和 ipv6 的时候有用
  -extscript string
        使用额外脚本来获取ipv4/ipv6地址。不使用留空
  -h    show this info
  -id int
        你的用户ID
  -install
        安装为服务运行 安装后就能用各自系统的服务管理工具进行管理 例如linux的systemctl
  -server
        是否以server方式免crontab方式运行
  -sub string
        sub子域名 eg: www 仅 type 为 ipv4/ipv4-local 和 ipv6 的时候有用 . 多个子域名用逗号隔开
  -token string
        你的token秘钥
  -type string
        使用类型：
        default 默认类型 使用公网ipv4地址替换web网站上配置的所有的域名信息
        ipv4  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前网络的公网IP
        ipv4-local  独立使用。更新 sub.domain 的域名的ipv4地址，默认情况下为使用当前设备的 内网ipv4地址 
        ipv6  独立使用。更新 sub.domain 的域名的ipv6地址，默认情况下为使用当前设备对外的IPV6地址
         (default "default")
  -uninstall
        卸载服务
```

## 关于 -extscript

写了个demo 在example文件夹下 [example](example)

```shell
#!/bin/sh
# 电信是 240e开头。 联通 2408 。 移动 2409。
pre_str="240e"
ip -6 address show |grep inet6 |awk '{print $2}' |grep "$pre_str" |grep -v '::' |cut -d '/' -f1
```

## 关于 deamon 后台运行 （由于需要注册服务，不能多个服务同时使用该模式）

> 需要多个服务免crontab配置运行，使用-server方式自行管理进程

Currently supports Windows XP+, Linux/(systemd | Upstart | SysV), and OSX/Launchd

example:

```shell
#安装
./ddns-go -id 1 -token 123456 -type ipv4 -sub test -domain domain.com -install
#启动
./ddns-go start
#停止
./ddns-go stop
#重启
./ddns-go restart
#卸载 
./ddns-go -uninstall
```

## docker实现

> docker run -d --net=host --name ddns1 -v /ddns-go:/usr/bin/ddns-go alpine /usr/bin/ddns-go -id 1 -token 123456 -type ipv4 -sub test -domain domain.com -server
>
或者 docker-compose的方式

```yaml
version: "3"
services:
  ddns:
    image: alpine
    restart: always
    network_mode: host
    volumes:
      - ./ddns-go:/usr/bin/ddns-go
    command:
      - /bin/sh
      - -c
      - /usr/bin/ddns-go  -id 1 -token 123456 -type ipv4 -sub test -domain domain.com -server
```