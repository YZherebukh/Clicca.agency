package blockchain

import (
	"reflect"
	"sync"

	"github.com/Clicca.agency/server/domain"
)

// BlockChain is a blockchain struct
type BlockChain struct {
	mu     *sync.RWMutex
	blocks []domain.Block
}

// New creates an empty blockchain
func New() *BlockChain {
	bc := &BlockChain{
		mu:     &sync.RWMutex{},
		blocks: make([]domain.Block, 0),
	}

	return bc
}

// Latest returns latest block in the chain
// if chain is empty, returns an empty block
func (bc *BlockChain) Latest() domain.Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if len(bc.blocks) > 0 {
		return bc.blocks[len(bc.blocks)-1]
	}

	return domain.Block{}
}

// Add is adding new block to the chain
// lock all operations with the chain
func (bc *BlockChain) Add(b domain.Block) bool {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	if len(bc.blocks) > 0 {
		prevHash := bc.blocks[len(bc.blocks)-1].Hash
		if !reflect.DeepEqual(prevHash, b.PrevHash) {
			return false
		}
	}

	bc.blocks = append(bc.blocks, b)
	return true
}

// All creates a copy of the chain and returns that copy
func (bc *BlockChain) All() []domain.Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	blocks := make([]domain.Block, len(bc.blocks))

	copy(blocks, bc.blocks)
	return blocks
}
