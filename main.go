package main

import (
	"log"
	"time"
)

func main() {
	// get config
	cfg, err := GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// get last known IP
	oldIp, err := GetLastIp()
	if err != nil {
		log.Fatalln(err)
	}

	// disable wireguard tunnel
	err = DisableService(oldIp)
	if err != nil {
		log.Fatalln(err)
	}

	// get new IP via DDNS
	newIp, err := GetIpFromDomain(cfg.Domain)
	if err != nil {
		log.Fatalln(err)
	}

	// set new host in wireguard
	err = SetWireguardHost(cfg.WireguardConfigFile, newIp)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	// enable new wireguard tunnel
	err = EnableService(newIp)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("success")
}
