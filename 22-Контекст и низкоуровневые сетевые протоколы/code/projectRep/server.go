package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()

		log.Printf("Получил: %s", text)

		if text == "quit" || text == "exit" {
			break
		}
		conn.Write(fmt.Appendf([]byte{}, fmt.Sprintf("Получил: %s\n", text)))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("Closing connection with %s", conn.RemoteAddr())
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:3302")

	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}

	log.Printf("Listen 0.0.0.0:3302")

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Cannot accept: %v", err)
		}
		go func() {
			defer conn.Close()
			handleConnection(conn)
		}()
	}
}
