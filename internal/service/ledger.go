package service

import (
	"errors"
	"time"

	"github.com/orobert/blockchain-supply-chain-integrity/internal/adapter/storage"
	"github.com/orobert/blockchain-supply-chain-integrity/internal/domain"
)

type LedgerService struct {
	repo storage.StorageAdapter
}

func NewLedgerService(repo storage.StorageAdapter) *LedgerService {
	return &LedgerService{repo: repo}
}

func (s *LedgerService) AddBlock(data [][]byte) (*domain.Block, error) {
	prevBlock, err := s.repo.GetLastBlock()
	if err != nil {
		return nil, err
	}

	index := 0
	prevHash := ""
	if prevBlock != nil {
		index = prevBlock.Index + 1
		prevHash = prevBlock.Hash
	}

	tree := domain.NewMerkleTree(data)
	if tree == nil {
		return nil, errors.New("no data to create merkle tree")
	}

	block := &domain.Block{
		Index:      index,
		Timestamp:  time.Now().Unix(),
		Data:       nil, // In a real system, data would be stored in the block or referenced
		PrevHash:   prevHash,
		MerkleRoot: tree.Hash,
	}

	block.Hash = block.CalculateHash()
	
	err = s.repo.SaveBlock(block)
	return block, err
}

func (s *LedgerService) ValidateChain() (bool, error) {
	lastBlock, err := s.repo.GetLastBlock()
	if err != nil {
		return false, err
	}
	if lastBlock == nil {
		return true, nil
	}

	for i := lastBlock.Index; i > 0; i-- {
		current, _ := s.repo.GetBlock(i)
		prev, _ := s.repo.GetBlock(i - 1)

		if current.Hash != current.CalculateHash() {
			return false, nil
		}
		if current.PrevHash != prev.Hash {
			return false, nil
		}
	}
	return true, nil
}
