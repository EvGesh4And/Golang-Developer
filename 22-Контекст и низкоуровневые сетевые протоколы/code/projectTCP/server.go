package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	localAddr := "0.0.0.0:52521"
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalf("Соединение открыть не удалось, видимо порт %s занят", localAddr)
	}
	log.Printf("Успешно начато прослушивание порта %s для входящих соединений", localAddr)
	defer listener.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			log.Println("Пришло время выключить сервер")
			listener.Close()
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Listener сломался: %v", err)
		}
		log.Printf("Установлено входящее соединение с адресса %s", conn.RemoteAddr())
		go func() {
			defer conn.Close()
			// handler(conn)
		}()
	}
	log.Println("Сервер завершил работу")
}
