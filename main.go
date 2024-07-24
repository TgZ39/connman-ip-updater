package main

import (
	"io"
	"log"
	"os"
	"time"
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
	log.Printf("loading config")
	cfg, err := GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// get last known IP
	log.Printf("reading last known IP")
	oldIp, err := GetLastIp(cfg.WireguardConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	// get new IP via DDNS
	log.Printf("looking up IP of %v", cfg.Domain)
	newIp, err := GetIpFromDomain(cfg.Domain)
	if err != nil {
		log.Fatalln(err)
	}

	// exit if DDNS hasn't updated
	if oldIp.Equal(newIp) {
		log.Printf("DDNS IP hasn't updated, exiting early")
		return
	}

	// disable Wireguard tunnel
	log.Printf("disabling connman service: %v", ToConnmanService(oldIp))
	err = DisableService(oldIp)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Second)

	// set new host in Wireguard
	log.Printf("updating host to %v in %v", newIp.String(), cfg.WireguardConfigFile)
	err = SetWireguardHost(cfg.WireguardConfigFile, newIp)
	if err != nil {
		log.Printf("error updating host: %v", err)

		log.Printf("enabling (broken) connman service")
		eErr := EnableService(oldIp)
		if eErr != nil {
			log.Printf("error enabling (broken) connman service: %v", eErr)
		}

		log.Fatalf("error setting new wireguard host: %v", err)
	}

	time.Sleep(time.Second)

	// enable new Wireguard tunnel
	log.Printf("enabling connman service: %v", ToConnmanService(newIp))
	err = EnableService(newIp)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("successfully updated wireguard tunnel")
}
