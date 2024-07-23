package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// set up logging
	logFile, err := getLogFile()
	if err != nil {
		log.Println(err)
	} else {
		mw := io.MultiWriter(logFile, os.Stdout)
		log.SetOutput(mw)
		defer logFile.Close()
	}

	// get config
	log.Println("loading config")
	cfg, err := GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// get last known IP
	log.Println("reading last known IP")
	oldIp, err := GetLastIp(cfg.WireguardConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	// get new IP via DDNS
	log.Println("looking up IP of", cfg.Domain)
	newIp, err := GetIpFromDomain(cfg.Domain)
	if err != nil {
		log.Fatalln(err)
	}

	// exit if DDNS hasn't updated
	if oldIp.Equal(newIp) {
		log.Println("DDNS IP hasn't updated, exiting early")
		return
	}

	// disable Wireguard tunnel
	log.Println("disabling connman service: ", ToConnmanService(oldIp))
	err = DisableService(oldIp)
	if err != nil {
		log.Fatalln(err)
	}

	// set new host in Wireguard
	log.Println("updating host to", newIp.String(), "in", cfg.WireguardConfigFile)
	err = SetWireguardHost(cfg.WireguardConfigFile, newIp)
	if err != nil {
		log.Println("error updating host: ", err)

		log.Println("enabling (broken) connman service")
		eErr := EnableService(oldIp)
		if eErr != nil {
			log.Println("error enabling (broken) connman service: ", eErr)
		}

		log.Fatalln(err)
	}

	// enable new Wireguard tunnel
	log.Println("enabling connman service:", ToConnmanService(newIp))
	err = EnableService(newIp)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("successfully updated wireguard tunnel")
}
