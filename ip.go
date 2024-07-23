package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

const lastIpFile = "last_ip.txt"

func GetLastIp() (net.IP, error) {
	buf, err := os.ReadFile(lastIpFile)
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(string(buf))
	if ip == nil {
		return nil, errors.New("file contains invalid IP")
	}

	return ip, nil
}

func SetLastIp(ip net.IP) error {
	return os.WriteFile(lastIpFile, []byte(ip.String()), 0644)
}

func ToConnmanService(ip net.IP) string {
	ipStr := ip.String()
	ipStr = strings.ReplaceAll(ipStr, ".", "_")

	return fmt.Sprintf("vpn_%v", ipStr)
}
