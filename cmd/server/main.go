package main

import (
	"log"
	"net"
	tm "twophasecommitservice/internal/tm"
	"twophasecommitservice/mocks"
	pb "twophasecommitservice/pkg/twophasecommit"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Create some mock services
	serviceA := &mocks.MockParticipant{Name: "ServiceA"}
	serviceB := &mocks.MockParticipant{Name: "ServiceB"}
	serviceC := &mocks.MockParticipant{Name: "ServiceC"}

	tm := tm.NewTransactionManager([]tm.Participant{serviceA, serviceB, serviceC})

	pb.RegisterTwoPhaseCommitServer(s, tm)

	log.Println("Serving gRPC on 0.0.0.0:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
