package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:3302")
	defer conn.Close()

	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	stdin := stdinScan()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		readRoutine(ctx, conn, wg)
		cancel()
	}()

	wg.Add(1)
	go func() {
		writeRoutine(ctx, conn, wg, stdin)
	}()

	wg.Wait()
}

func writeRoutine(ctx context.Context, conn net.Conn, wg *sync.WaitGroup, stdin chan string) {
	defer wg.Done()
	defer log.Printf("Finished writeRoutine")

	for {
		select {
		case <-ctx.Done():
			return
		case str := <-stdin:
			_ = str
			log.Printf("To server %s", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}
	}
}

func readRoutine(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer log.Printf("Finished readRoutine")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		_ = text
		log.Printf("Read: %v", text)
	}
}

func stdinScan() chan string {
	out := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if scanner.Err() != nil {
			close(out)
		}
	}()
	return out
}
