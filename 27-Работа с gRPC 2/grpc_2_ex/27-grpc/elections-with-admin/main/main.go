package main

import (
	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-admin/pb"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			StreamServerRequestValidatorInterceptor(ValidateReq),
		),
	)
	pb.RegisterElectionsServer(grpcServer, NewService())
	reflection.Register(grpcServer) // postman

	log.Printf("starting grpcServer on %s", lsn.Addr().String())
	if err := grpcServer.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
