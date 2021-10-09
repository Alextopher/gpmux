package main

import (
	"fmt"
	"log"
	"net"
)

func connect(host string, port uint16) *net.UDPConn {
	raddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalln("Failed to resolve addr with err:", err)
	}

	laddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		log.Fatalln("Failed to resolve addr with err:", err)
	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		log.Fatalln("Failed to connect with err:", err)
	}
	return conn
}
