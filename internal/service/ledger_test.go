package service

import (
	"testing"
	"github.com/orobert/blockchain-supply-chain-integrity/internal/domain"
)

type MockStorage struct {
	blocks []*domain.Block
}

func (m *MockStorage) SaveBlock(block *domain.Block) error {
	m.blocks = append(m.blocks, block)
	return nil
}

func (m *MockStorage) GetBlock(index int) (*domain.Block, error) {
	if index < 0 || index >= len(m.blocks) {
		return nil, nil
	}
	return m.blocks[index], nil
}

func (m *MockStorage) GetLastBlock() (*domain.Block, error) {
	if len(m.blocks) == 0 {
		return nil, nil
	}
	return m.blocks[len(m.blocks)-1]
}

func TestLedgerIntegration(t *testing.T) {
	mock := &MockStorage{}
	svc := NewLedgerService(mock)

	// Add first block
	data1 := [][]byte{[]byte("item1_origin"), []byte("item1_qc_passed")}
	b1, err := svc.AddBlock(data1)
	if err != nil {
		t.Fatalf("Failed to add block 1: %v", err)
	}

	if b1.Index != 0 {
		t.Errorf("Expected index 0, got %d", b1.Index)
	}

	// Add second block
	data2 := [][]byte{[]byte("item1_shipped")}
	b2, err := svc.AddBlock(data2)
	if err != nil {
		t.Fatalf("Failed to add block 2: %v", err)
	}

	if b2.PrevHash != b1.Hash {
		t.Error("Block 2 should point to block 1")
	}

	// Validate chain
	valid, err := svc.ValidateChain()
	if err != nil {
		t.Fatalf("Validation error: %v", err)
	}
	if !valid {
		t.Error("Chain should be valid")
	}

	// Tamper and re-validate
	b1.Hash = "corrupted"
	valid, _ = svc.ValidateChain()
	if valid {
		t.Error("Chain should be invalid after tampering")
	}
}
