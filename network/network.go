// Package network is used to find network information in the local environment
package network

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal("Error is: ", err)
	}

	netInterfaceMatch := false

	for _, addr := range addrs {
		ipv6Addr, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			log.Fatal("Error is: ", err)
		}

		if localAddr.IP.String() == ipv6Addr.String() {
			netInterfaceMatch = true
		}
	}

	if !netInterfaceMatch {
		fmt.Println("Cannot determine your IP address.")
		os.Exit(1)
	}

	return localAddr.IP
}
