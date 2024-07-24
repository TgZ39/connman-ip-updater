package main

import (
	"net"
	"slices"
)

func IsWireguardActive() (bool, error) {
	wgInterfaces := []string{"wg0", "wg1", "wg2", "wg3", "wg4", "wg5"}

	interfaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	for _, inf := range interfaces {
		if slices.Contains(wgInterfaces, inf.Name) {
			return true, nil
		}
	}

	return false, nil
}
