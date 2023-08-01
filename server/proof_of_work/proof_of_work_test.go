package proof_of_work_test

import (
	"testing"

	"github.com/Clicca.agency/server/proof_of_work"
	"github.com/stretchr/testify/assert"
)

func TestGetTask(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		pow := proof_of_work.New(nil, 1)

		task := pow.GetTask()

		assert.Equal(t, 256, task.BitLen())
	})
}
