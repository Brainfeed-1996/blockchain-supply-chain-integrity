package domain

import (
	"crypto/sha256"
	"fmt"
)

// ZKPProof represents a mock Zero-Knowledge Proof for supply chain verification.
type ZKPProof struct {
	Proof     []byte
	PublicInput []byte
}

func (p *ZKPProof) Verify() bool {
	// Mock ZKP verification logic. 
	// In a real implementation, this would use a library like gnark or bellman.
	// Logic: If the SHA256 of the proof matches a specific property derived from PublicInput.
	h := sha256.Sum256(p.Proof)
	return fmt.Sprintf("%x", h[:])[0:4] == "0000" // Simplified proof-of-work style ZKP mock
}

type Block struct {
	Index        int
	Timestamp    int64
	Data         []byte
	PrevHash     string
	Hash         string
	Signature    []byte
	Validator    string
	MerkleRoot   string
	ZKPProof     *ZKPProof // Integrated ZKP
}

func (b *Block) CalculateHash() string {
	record := fmt.Sprintf("%d%d%x%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash, b.MerkleRoot)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  string
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := &MerkleNode{}
	if left == nil && right == nil {
		h := sha256.Sum256(data)
		node.Hash = fmt.Sprintf("%x", h[:])
	} else {
		prevHashes := left.Hash + right.Hash
		h := sha256.Sum256([]byte(prevHashes))
		node.Hash = fmt.Sprintf("%x", h[:])
	}
	node.Left = left
	node.Right = right
	return node
}

func NewMerkleTree(data [][]byte) *MerkleNode {
	var nodes []*MerkleNode
	for _, d := range data {
		nodes = append(nodes, NewMerkleNode(nil, nil, d))
	}
	if len(nodes) == 0 {
		return nil
	}
	for len(nodes) > 1 {
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		var newLevel []*MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			node := NewMerkleNode(nodes[i], nodes[i+1], nil)
			newLevel = append(newLevel, node)
		}
		nodes = newLevel
	}
	return nodes[0]
}
