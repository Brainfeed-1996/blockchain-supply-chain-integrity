package domain

import (
	"testing"
)

func TestMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("tx1"),
		[]byte("tx2"),
		[]byte("tx3"),
	}

	tree := NewMerkleTree(data)
	if tree == nil {
		t.Fatal("Merkle tree should not be nil")
	}

	if tree.Hash == "" {
		t.Error("Merkle root hash should not be empty")
	}
}

func TestBlockValidation(t *testing.T) {
	block := &Block{
		Index:      1,
		Timestamp:  123456789,
		Data:       []byte("test"),
		PrevHash:   "prev",
		MerkleRoot: "merkle",
	}

	hash := block.CalculateHash()
	block.Hash = hash

	if block.Hash != block.CalculateHash() {
		t.Error("Block hash should be valid")
	}

	block.Data = []byte("tampered")
	if block.Hash == block.CalculateHash() {
		t.Error("Block hash should change when data is tampered")
	}
}
