// Package ip 地址相关的工具类
package ip

import (
	"encoding/binary"
	"net"
)

// ToLong 将IPv4协议的IP地址转换为整数，当转换成功时ok返回true，rs返回具体的数据
func ToLong(ip string) (rs uint32, ok bool) {
	if ip == "" {
		return
	}
	rs = binary.BigEndian.Uint32(net.ParseIP(ip).To16()[12:16])
	if rs > 0 {
		ok = true
	}
	return
}

// LongToIP 将整数转换为IPv4协议的ip地址
func LongToIP(long uint32) string {
	var newIP net.IP
	newIP = make(net.IP, net.IPv4len)
	binary.BigEndian.PutUint32(newIP, long)
	return newIP.String()
}
