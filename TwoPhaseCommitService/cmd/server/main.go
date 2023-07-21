package main

import (
	tm "TwoPhaseCommitService/TwoPhaseCommitService/internal/tm"
	pb "TwoPhaseCommitService/TwoPhaseCommitService/pkg/twophasecommit"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	tm := tm.NewTransactionManager()
	pb.RegisterTwoPhaseCommitServer(s, tm)

	log.Println("Serving gRPC on 0.0.0.0:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
