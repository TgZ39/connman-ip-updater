package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

func GetLastIp(configFile string) (net.IP, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "Host = ") {
			ipStr := strings.TrimSpace(text)
			ipStr = strings.TrimPrefix(ipStr, "Host = ")

			ip := net.ParseIP(ipStr)
			if ip == nil {
				return nil, errors.New("invalid IP in wireguard config")
			}

			return ip, nil
		}
	}

	return nil, errors.New("couldn't find \"Host = x.x.x.x\" in wireguard config")
}

// ToConnmanService returns given IP as connman service, e.g. 192.168.178.1 -> vpn_192_168_178_1
func ToConnmanService(ip net.IP) string {
	ipStr := ip.String()
	ipStr = strings.ReplaceAll(ipStr, ".", "_")

	return fmt.Sprintf("vpn_%v", ipStr)
}
