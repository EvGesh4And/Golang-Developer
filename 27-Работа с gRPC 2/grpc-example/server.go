package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-example/proto" // Импорт сгенерированного кода
)

// Реализация сервиса HelloService
type helloServer struct {
	proto.UnimplementedHelloServiceServer
}

// Метод SayHello
func (s *helloServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	message := fmt.Sprintf("Привет, %s!", req.GetName())
	return &proto.HelloReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("не удалось прослушать: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &helloServer{})

	log.Println("gRPC-сервер слушает на :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}
}
