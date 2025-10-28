package app

import (
	"fmt"
	"net"
	"time"
)

func PingOne(domain string, port string) {

	address := domain + ":" + port
	timeout := time.Duration(2 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Printf("[DOWN] %v is unreachable, \nError: %v", domain, err)
	} else {
		fmt.Printf("[UP] %v is reachable, \nFrom: %v\nTo: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
	}
}
