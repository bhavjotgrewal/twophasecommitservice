package tm

import (
	pb "TwoPhaseCommitService/TwoPhaseCommitService/pkg/twophasecommit"
	"context"
)

type transactionManager struct {
	pb.UnimplementedTwoPhaseCommitServer
}

func NewTransactionManager() *transactionManager {
	return &transactionManager{}
}

func (tm *transactionManager) Prepare(ctx context.Context, in *pb.PrepareRequest) (*pb.VoteResponse, error) {
	// TODO: Implement prepare phase
	return &pb.VoteResponse{}, nil
}

func (tm *transactionManager) Commit(ctx context.Context, in *pb.CommitRequest) (*pb.AckResponse, error) {
	// TODO: Implement commit phase
	return &pb.AckResponse{}, nil
}
