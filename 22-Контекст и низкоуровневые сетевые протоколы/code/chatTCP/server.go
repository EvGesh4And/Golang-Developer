package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Fprintf(conn, "Welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("RECEIVED: %v", text)
		if text == "quit" || text == "exit" {
			break
		}

		conn.Write([]byte(fmt.Sprintf("I have received %s\n", text)))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("Closing connection with %s", conn.RemoteAddr())
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:4242")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}

	log.Printf("Listen 127.0.0.1:4242")

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
