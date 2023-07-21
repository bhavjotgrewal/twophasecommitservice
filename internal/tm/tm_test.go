package tm

import (
	"context"
	"testing"
	"twophasecommitservice/mocks"
	pb "twophasecommitservice/pkg/twophasecommit"
)

func TestTransactionManager_Prepare(t *testing.T) {
	participants := []Participant{
		&mocks.MockParticipant{Name: "participant1", PrepareSuccess: true},
		&mocks.MockParticipant{Name: "participant2", PrepareSuccess: true},
		&mocks.MockParticipant{Name: "participant3", PrepareSuccess: true},
	}

	tm := NewTransactionManager(participants)

	// Prepare the transaction
	res, err := tm.Prepare(context.Background(), &pb.PrepareRequest{TransactionId: "123"})
	if err != nil {
		t.Fatalf("unexpected error from Prepare: %v", err)
	}

	if res.Vote != pb.VoteResponse_YES {
		t.Errorf("expected vote YES, got %s", res.Vote)
	}

	// Check that all participants have been prepared
	for _, participant := range participants {
		if p := participant.(*mocks.MockParticipant); !p.Prepared {
			t.Errorf("expected participant %s to be prepared", p.Name)
		}
	}
}

func TestTransactionManager_Commit(t *testing.T) {
	participants := []Participant{
		&mocks.MockParticipant{Name: "participant1", PrepareSuccess: true, CommitSuccess: true},
		&mocks.MockParticipant{Name: "participant2", PrepareSuccess: true, CommitSuccess: true},
		&mocks.MockParticipant{Name: "participant3", PrepareSuccess: true, CommitSuccess: true},
	}

	tm := NewTransactionManager(participants)

	// Prepare the transaction
	_, err := tm.Prepare(context.Background(), &pb.PrepareRequest{TransactionId: "123"})
	if err != nil {
		t.Fatalf("unexpected error from Prepare: %v", err)
	}

	// Commit the transaction
	res, err := tm.Commit(context.Background(), &pb.CommitRequest{TransactionId: "123"})
	if err != nil {
		t.Fatalf("unexpected error from Commit: %v", err)
	}

	if !res.Success {
		t.Errorf("expected commit to be successful")
	}

	// Check that all participants have been committed
	for _, participant := range participants {
		if p := participant.(*mocks.MockParticipant); !p.Committed {
			t.Errorf("expected participant %s to be committed", p.Name)
		}
	}
}

func TestTransactionManager_PrepareFails(t *testing.T) {
	participants := []Participant{
		&mocks.MockParticipant{Name: "participant1", PrepareSuccess: true},
		&mocks.MockParticipant{Name: "participant2", PrepareSuccess: false},
		&mocks.MockParticipant{Name: "participant3", PrepareSuccess: true},
	}

	tm := NewTransactionManager(participants)

	// Prepare the transaction
	res, err := tm.Prepare(context.Background(), &pb.PrepareRequest{TransactionId: "123"})
	if err != nil {
		t.Fatalf("unexpected error from Prepare: %v", err)
	}

	if res.Vote != pb.VoteResponse_NO {
		t.Errorf("expected vote NO, got %s", res.Vote)
	}
}

func TestTransactionManager_CommitFails(t *testing.T) {
	participants := []Participant{
		&mocks.MockParticipant{Name: "participant1", PrepareSuccess: true, CommitSuccess: true},
		&mocks.MockParticipant{Name: "participant2", PrepareSuccess: true, CommitSuccess: false},
		&mocks.MockParticipant{Name: "participant3", PrepareSuccess: true, CommitSuccess: true},
	}

	tm := NewTransactionManager(participants)

	// Prepare the transaction
	_, err := tm.Prepare(context.Background(), &pb.PrepareRequest{TransactionId: "123"})
	if err != nil {
		t.Fatalf("unexpected error from Prepare: %v", err)
	}

	// Commit the transaction
	res, err := tm.Commit(context.Background(), &pb.CommitRequest{TransactionId: "123"})
	if err != nil {
		t.Fatalf("unexpected error from Commit: %v", err)
	}

	if res.Success {
		t.Errorf("expected commit to fail")
	}
}
