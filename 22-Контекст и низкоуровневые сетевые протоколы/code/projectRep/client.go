package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func stdin(ctx context.Context) chan string {
	out := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			case out <- scanner.Text():
			}
		}
	}()
	return out
}

func read() {
	
}

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)

	chStdin := stdin(ctx)

	conn, err := net.Dial("tcp", "0.0.0.0:3302")

	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}

	log.Println("Подключились к 0.0.0.0:3302")
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for {
		select {
		case <-ctx.Done():
			log.Println("Завершение сеанса")
			break
		case text := <-chStdin:
			conn.Write(fmt.Appendf([]byte{}, "%s\n", text))
			scanner.Scan()
			log.Println(scanner.Text())
		}
	}

	log.Println("Завершение сеанса")
}
