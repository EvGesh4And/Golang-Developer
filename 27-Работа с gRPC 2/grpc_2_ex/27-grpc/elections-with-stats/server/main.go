package main

import (
	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-stats/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterElectionsServer(server, NewService())

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
