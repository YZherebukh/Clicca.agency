package task

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"

	"github.com/Clicca.agency/worker/domain"
)

type BlockChain struct {
	difficulty int
}

func NewBlockChain(difficulty int) *BlockChain {
	bc := &BlockChain{
		difficulty: difficulty,
	}

	return bc
}

func (bc *BlockChain) Resolve(block domain.Block, target *big.Int) (*domain.Block, error) {
	var (
		intHash big.Int
		hash    [32]byte
		nonce   int
	)

	diff, err := toHex(int64(bc.difficulty))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve task, %w", err)
	}

	for nonce < math.MaxInt64 {
		non, err := toHex(int64(nonce))
		if err != nil {
			return nil, fmt.Errorf("failed to resolve task, %w", err)
		}

		data := bytes.Join(
			[][]byte{
				block.PrevHash,
				block.Data,
				diff,
				non,
			},
			[]byte{},
		)

		hash = sha256.Sum256(data)

		intHash.SetBytes(hash[:])

		if intHash.Cmp(target) == -1 {
			return &domain.Block{
				Hash:     hash[:],
				Data:     block.Data,
				PrevHash: block.PrevHash,
				Nonce:    nonce,
			}, nil
		}

		nonce++
	}

	return nil, fmt.Errorf("hash not found")
}

func toHex(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, fmt.Errorf("toHex failed: %w", err)
	}
	return buff.Bytes(), nil
}
