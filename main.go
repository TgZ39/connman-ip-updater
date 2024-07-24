package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// set up logging
	lf, err := getLogFile()
	if err != nil {
		log.Printf("error setting up log file: %v", err)
	} else {
		mw := io.MultiWriter(lf, os.Stdout)
		log.SetOutput(mw)
		defer lf.Close()
	}

	wgActive, err := IsWireguardActive()
	if err != nil {
		log.Fatalf("could determine whether a wireguard interface is active: %v", err)
	}

	// get config
	log.Printf("loading config")
	cfg, err := GetConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// get last known IP
	log.Printf("reading last known IP")
	oldIp, err := GetLastIp(cfg.WireguardConfigFile)
	if err != nil {
		log.Fatalf("error reading last known IP: %v", err)
	}

	// get new IP via DDNS
	log.Printf("looking up IP of %v", cfg.Domain)
	newIp, err := GetIpFromDomain(cfg.Domain)
	if err != nil {
		log.Fatalf("error looking up IP: %v", err)
	}

	// exit if DDNS hasn't updated
	if oldIp.Equal(newIp) {
		if !wgActive && cfg.Enabled {
			// enable new Wireguard tunnel
			log.Printf("enabling connman service: %v", ToConnmanService(newIp))
			err = EnableService(newIp)
			if err != nil {
				log.Fatalf("error enabling wireguard tunnel: %v", err)
			}
		} else if wgActive && !cfg.Enabled {
			// disable Wireguard tunnel
			log.Printf("disabling connman service: %v", ToConnmanService(oldIp))
			err = DisableService(oldIp)
			if err != nil {
				log.Fatalf("error disabling wireguard tunnel: %v", err)
			}

			time.Sleep(time.Second)
		}

		log.Printf("DDNS IP hasn't updated, exiting early")
		return
	}

	if wgActive {
		// disable Wireguard tunnel
		log.Printf("disabling connman service: %v", ToConnmanService(oldIp))
		err = DisableService(oldIp)
		if err != nil {
			log.Fatalf("error disabling wireguard tunnel: %v", err)
		}

		time.Sleep(time.Second)
	}

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

	if cfg.Enabled {
		time.Sleep(time.Second)

		// enable new Wireguard tunnel
		log.Printf("enabling connman service: %v", ToConnmanService(newIp))
		err = EnableService(newIp)
		if err != nil {
			log.Fatalf("error enabling wireguard tunnel: %v", err)
		}
	}

	log.Printf("successfully updated wireguard tunnel")
}
