package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func DisableService(ip net.IP) error {
	service := ToConnmanService(ip)

	cmd := exec.Command("connmanctl", "disconnect", service)
	err := cmd.Run()
	return err
}

func EnableService(ip net.IP) error {
	service := ToConnmanService(ip)

	cmd := exec.Command("connmanctl", "connect", service)
	err := cmd.Run()
	return err
}

func SetWireguardHost(configFile string, ip net.IP) error {
	file, err := os.Open(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	foundHost := false
	buf := bytes.Buffer{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "Host = ") {
			newLine := fmt.Sprintf("Host = %v\n", ip.String())
			buf.WriteString(newLine)
			foundHost = true
		} else {
			buf.WriteString(fmt.Sprintln(text))
		}
	}

	if !foundHost {
		return errors.New("host line in wireguard config wasn't found: \"Host = ...\"")
	}

	err = os.WriteFile(configFile, buf.Bytes(), 0664)
	if err != nil {
		return err
	}

	return nil
}
