package main

import (
	"context"
	"fmt"
	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-stats/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewElectionsClient(conn)

	stream, errStat := client.GetStats(context.Background(), &emptypb.Empty{})
	if errStat != nil {
		log.Fatal(errStat)
	}

	for i := 0; i < 5; i++ {
		response, errRecv := stream.Recv()
		if errRecv != nil {
			log.Fatal(errRecv)
		}

		fmt.Println(response)
		time.Sleep(1 * time.Second)
	}

	errClose := stream.CloseSend()
	if errClose != nil {
		log.Fatal(errClose)
	}
}
