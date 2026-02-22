package storage

import "github.com/orobert/blockchain-supply-chain-integrity/internal/domain"

type StorageAdapter interface {
	SaveBlock(block *domain.Block) error
	GetBlock(index int) (*domain.Block, error)
	GetLastBlock() (*domain.Block, error)
}

type MemoryStorage struct {
	blocks []*domain.Block
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{blocks: make([]*domain.Block, 0)}
}

func (s *MemoryStorage) SaveBlock(block *domain.Block) error {
	s.blocks = append(s.blocks, block)
	return nil
}

func (s *MemoryStorage) GetBlock(index int) (*domain.Block, error) {
	if index < 0 || index >= len(s.blocks) {
		return nil, nil
	}
	return s.blocks[index], nil
}

func (s *MemoryStorage) GetLastBlock() (*domain.Block, error) {
	if len(s.blocks) == 0 {
		return nil, nil
	}
	return s.blocks[len(s.blocks)-1], nil
}
