package blockchain_test

import (
	"testing"

	"github.com/Clicca.agency/server/domain"
	blockchain "github.com/Clicca.agency/server/internal/block_chain"
	"github.com/stretchr/testify/assert"
)

func TestLatest(t *testing.T) {
	t.Run("empty_chain", func(t *testing.T) {
		testBlockChain := blockchain.New()

		latest := testBlockChain.Latest()
		assert.Equal(t, domain.Block{}, latest)
	})
	t.Run("not_empty_chain", func(t *testing.T) {
		testBlockChain := blockchain.New()
		testBlockChain.Add(domain.Block{Data: []byte("test_data")})

		latest := testBlockChain.Latest()
		assert.Equal(t, []byte("test_data"), latest.Data)
	})
}

func TestAdd(t *testing.T) {
	t.Run("empty_chain", func(t *testing.T) {
		testBlockChain := blockchain.New()

		latest := testBlockChain.Latest()
		assert.True(t, len(latest.PrevHash) == 0)
	})
	t.Run("not_empty_chain", func(t *testing.T) {
		testBlockChain := blockchain.New()
		testBlockChain.Add(domain.Block{Data: []byte("test_data_1"), Hash: []byte("test_hash")})
		testBlockChain.Add(domain.Block{Data: []byte("test_data_2"), Hash: []byte("test_hash")})

		latest := testBlockChain.Latest()

		assert.Equal(t, []byte("test_hash"), latest.PrevHash)
	})
}

func TestAll(t *testing.T) {
	t.Run("empty_chain", func(t *testing.T) {
		testBlockChain := blockchain.New()

		chain := testBlockChain.All()
		assert.True(t, len(chain) == 0)
	})
	t.Run("not_empty_chain", func(t *testing.T) {
		testBlockChain := blockchain.New()
		testBlockChain.Add(domain.Block{Data: []byte("test_data_1"), Hash: []byte("test_hash")})
		testBlockChain.Add(domain.Block{Data: []byte("test_data_2"), Hash: []byte("test_hash")})

		chain := testBlockChain.All()
		assert.True(t, len(chain) == 2)
	})
}
