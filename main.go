package main

import (
	"log"
	"time"
)

func main() {
	// get config
	log.Println("loading config: ip_updater_config.toml")
	cfg, err := GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// get last known IP
	log.Println("reading last_ip.txt")
	oldIp, err := GetLastIp()
	if err != nil {
		log.Fatalln(err)
	}

	// get new IP via DDNS
	log.Println("looking up IPv4 of", cfg.Domain)
	newIp, err := GetIpFromDomain(cfg.Domain)
	if err != nil {
		log.Fatalln(err)
	}

	// exit if DDNS hasn't updated
	if oldIp.Equal(newIp) {
		log.Println("DDNS ip hasn't updated, exiting early")
		return
	}

	// disable Wireguard tunnel
	log.Println("disabling connman service: ", ToConnmanService(oldIp))
	err = DisableService(oldIp)
	if err != nil {
		log.Fatalln(err)
	}

	// set new host in Wireguard
	log.Println("updating host to ", newIp.String(), " in ", cfg.WireguardConfigFile)
	err = SetWireguardHost(cfg.WireguardConfigFile, newIp)
	if err != nil {
		log.Println("error updating host")

		log.Println("enabling (broken) connman service")
		eErr := EnableService(oldIp)
		if eErr != nil {
			log.Println("error enabling (broken) connman service: ", err)
		}

		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	// enable new Wireguard tunnel
	log.Println("enabling connman service: ", ToConnmanService(newIp))
	err = EnableService(newIp)
	if err != nil {
		log.Fatalln(err)
	}

	// set new IP in last_ip.txt
	log.Println("setting last_ip.txt to ", newIp.String())
	err = SetLastIp(newIp)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("successfully updated wireguard tunnel")
}
