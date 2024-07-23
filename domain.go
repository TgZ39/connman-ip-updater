package main

import (
	"net"
)

func GetIpFromDomain(domain string) (net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	return ips[0], nil
}
