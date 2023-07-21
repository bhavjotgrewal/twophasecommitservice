package mocks

import "log"

type MockParticipant struct {
	Name           string
	Prepared       bool
	Committed      bool
	PrepareSuccess bool
	CommitSuccess  bool
}

func (m *MockParticipant) Prepare(txnID string) bool {
	log.Printf("%s is preparing transaction %s\n", m.Name, txnID)
	if m.PrepareSuccess {
		m.Prepared = true
	}
	return m.PrepareSuccess
}

func (m *MockParticipant) Commit(txnID string) bool {
	log.Printf("%s is committing transaction %s\n", m.Name, txnID)
	if m.CommitSuccess {
		m.Committed = true
	}
	return m.CommitSuccess
}
