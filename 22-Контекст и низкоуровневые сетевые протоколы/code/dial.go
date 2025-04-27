package main

import (
	"fmt"
	"net"
)

func main() {
	Conn, err := net.Dial("tcp", "golang.org:http")

	fmt.Println(Conn, err)
	b := make([]byte, 0, 1000)
	Conn.Read(b)
	fmt.Println(b)
}
