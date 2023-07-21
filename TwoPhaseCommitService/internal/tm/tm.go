package tm

import (
	pb "TwoPhaseCommitService/TwoPhaseCommitService/pkg/twophasecommit"
	"context"
	"sync"
)

type Participant interface {
	Prepare(txnID string) bool
	Commit(txnID string) bool
}

type transactionManager struct {
	pb.UnimplementedTwoPhaseCommitServer
	participants []Participant
	mux          sync.Mutex
	prepared     map[string]bool
}

func NewTransactionManager(participants []Participant) *transactionManager {
	return &transactionManager{
		participants: participants,
		prepared:     make(map[string]bool),
	}
}

func (tm *transactionManager) Prepare(ctx context.Context, in *pb.PrepareRequest) (*pb.VoteResponse, error) {
	txnID := in.GetTransactionId()
	vote := pb.VoteResponse_YES
	tm.mux.Lock()
	defer tm.mux.Unlock()

	for _, participant := range tm.participants {
		if !participant.Prepare(txnID) {
			// If any participant votes NO, mark the transaction as not prepared and return immediately
			tm.prepared[txnID] = false
			return &pb.VoteResponse{Vote: pb.VoteResponse_NO}, nil
		}
	}

	tm.prepared[txnID] = true
	return &pb.VoteResponse{Vote: vote}, nil
}

func (tm *transactionManager) Commit(ctx context.Context, in *pb.CommitRequest) (*pb.AckResponse, error) {
	txnID := in.GetTransactionId()
	success := false
	tm.mux.Lock()
	defer tm.mux.Unlock()

	if tm.prepared[txnID] {
		success = true
		for _, participant := range tm.participants {
			if !participant.Commit(txnID) {
				success = false
				break
			}
		}
	}

	// If any participant's commit failed, mark the transaction as not prepared
	if !success {
		tm.prepared[txnID] = false
	}

	return &pb.AckResponse{Success: success}, nil
}
