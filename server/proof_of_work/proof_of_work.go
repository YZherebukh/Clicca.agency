package proof_of_work

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/Clicca.agency/server/domain"
)

// Chainer is a blockChain interface
type Chainer interface {
	Latest() domain.Block
}

// ProofOfWork is a proof of work struct
type ProofOfWork struct {
	chainer    Chainer
	Target     *big.Int
	difficulty int
}

// new creates new proof of work instance
func New(chainer Chainer, difficulty int) *ProofOfWork {
	return &ProofOfWork{
		chainer:    chainer,
		difficulty: difficulty,
		Target:     big.NewInt(1),
	}
}

// GetTask sets and returnes a task for POW
func (pow ProofOfWork) GetTask() *big.Int {
	target := big.NewInt(1)
	pow.Target.Lsh(target, uint(256-pow.difficulty))

	return pow.Target
}

// Validate is taking a block and comapers if task was resolved
func (pow ProofOfWork) Validate(block domain.Block) (bool, error) {
	var intHash big.Int

	if intHash.SetBytes(block.PrevHash).Cmp(intHash.SetBytes(pow.chainer.Latest().Hash)) != 0 {
		return false, fmt.Errorf("previous hash is does not match")
	}

	hexDifficulty, err := toHex(int64(pow.difficulty))
	if err != nil {
		return false, fmt.Errorf("failed to Hex, error: %w", err)
	}
	hexNonce, err := toHex(int64(block.Nonce))
	if err != nil {
		return false, fmt.Errorf("failed to Hex, error: %w", err)
	}

	data := bytes.Join(
		[][]byte{
			block.PrevHash,
			block.Data,
			hexDifficulty,
			hexNonce,
		},
		[]byte{},
	)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1, nil
}

func toHex(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
