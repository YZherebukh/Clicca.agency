package executor

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/Clicca.agency/worker/domain"
)

type Client interface {
	GetTask(ctx context.Context) (*domain.TaskResponse, error)
	VerifyTask(ctx context.Context, block domain.Block) (*domain.VerifyResponse, error)
}

type Tasker interface {
	Resolve(block domain.Block, target *big.Int) (*domain.Block, error)
}

type Executor struct {
	c Client
	t Tasker
}

func New(c Client, t Tasker) *Executor {
	return &Executor{
		c: c,
		t: t,
	}
}

func (e *Executor) Do(ctx context.Context) error {
	taskResp, err := e.c.GetTask(ctx)
	if err != nil {
		return fmt.Errorf("executor failed %w", err)
	}

	block, err := e.t.Resolve(*taskResp.Block, taskResp.Task)
	if err != nil {
		return fmt.Errorf("executor failed %w", err)
	}

	verResp, err := e.c.VerifyTask(ctx, *block)
	if err != nil {
		return fmt.Errorf("executor failed %w", err)
	}

	log.Printf("success, here's your quote:  \"%s\" \n", verResp.Quote)
	return nil
}
