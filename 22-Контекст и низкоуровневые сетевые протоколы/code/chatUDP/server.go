package main

import (
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3303")
	if err != nil {
		log.Fatalf("Error addr 0.0.0.0:3303")
	}

	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Cannot listen 0.0.0.0:3303")
	}
	defer l.Close()
	log.Printf("listen 0.0.0.0:3303")
	msg := make([]byte, 10000)

	for {
		length, fromAddr, err := l.ReadFromUDP(msg)
		if err != nil {
			log.Fatalf("Error happened")
		}

		log.Printf("Message from %s with length %d: %s", fromAddr.String(), length, string(msg))
	}
}
