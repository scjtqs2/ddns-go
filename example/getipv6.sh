#!/bin/bash
# 电信是 240e开头。 联通 2408 。 移动 2409。
pre_str="240e"
ip -6 address show |grep inet6 |awk '{print $2}' |grep "$pre_str" |grep -v '::' |cut -d '/' -f1