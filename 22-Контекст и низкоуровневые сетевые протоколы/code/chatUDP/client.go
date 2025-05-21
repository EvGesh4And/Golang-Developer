package main

import (
	"log"
	"net"
	"time"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3303")
	if err != nil {
		log.Fatalf("cannot resolve server addr: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("cannot dial to server: %v", err)
	}
	defer conn.Close()
	log.Printf("write 0.0.0.0:3303")
	order := 0
	for {
		msg := ""
		switch order {
		case 0:
			msg = "1"
			order++
		case 1:
			msg = "2"
			order++
		case 2:
			msg = "3"
			order = 0
		}
		log.Printf("message: %s sent", msg)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatalf("cannot send: %v", err)
		}
		time.Sleep(time.Second)
	}
}
