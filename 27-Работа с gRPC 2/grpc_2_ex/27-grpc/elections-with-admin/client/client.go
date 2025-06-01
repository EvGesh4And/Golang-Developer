package main

import (
	"context"
	"errors"
	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections-with-admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewElectionsClient(conn)

	stream, errInternal := client.Internal(context.Background())
	if errInternal != nil {
		log.Fatal(errInternal)
	}

	done := make(chan struct{})

	go func() {
		for candidateID := uint32(0); candidateID < 5; candidateID++ {
			vote := &pb.Vote{
				Passport:    "100",
				CandidateId: candidateID,
			}

			if errSend := stream.Send(vote); errSend != nil {
				log.Fatal(errSend)
			}

			log.Printf("1: vote submitted\n")
			time.Sleep(1500 * time.Millisecond)
		}

		if err := stream.CloseSend(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			log.Printf("2: BODY received: %s; STATS received: %s; VOTE received: %s\n", resp.GetBody(), resp.GetStats(), resp.GetVote())
		}
	}()

	go func() {
		<-stream.Context().Done()
		if err := stream.Context().Err(); err != nil {
			log.Println(err, errors.Is(err, context.Canceled))
		}
		done <- struct{}{}
		close(done)
	}()

	<-done

	log.Println("3: client finished")
}
