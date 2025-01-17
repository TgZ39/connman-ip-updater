package main

import (
	"log"
	"net"
	"time"
)

func GetIpFromDomain(domain string) (net.IP, error) {
	var ips []net.IP
	var err error

	for i := 0; i < 5; i++ {
		ips, err = net.LookupIP(domain)
		if err != nil {
			log.Printf("error looking up IP: %v", err)
			log.Printf("trying again in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			return ips[0], err
		}
	}

	return nil, err
}
