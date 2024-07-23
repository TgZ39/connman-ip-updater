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

	ipStr := string(buf)
	ipStr = strings.TrimSpace(ipStr)

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, errors.New(fmt.Sprintf("file contains invalid IP: %v", string(buf)))
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
