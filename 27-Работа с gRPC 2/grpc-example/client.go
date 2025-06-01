package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"grpc-example/proto" // Импорт сгенерированного кода
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Мир"})
	if err != nil {
		log.Fatalf("ошибка при вызове SayHello: %v", err)
	}

	log.Printf("Ответ от сервера: %s", resp.GetMessage())
}
