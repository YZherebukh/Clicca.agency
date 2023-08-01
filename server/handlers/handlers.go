package handlers

import (
	"math/big"
	"net/http"

	"github.com/Clicca.agency/server/domain"
	"github.com/monzo/typhon"
)

// Quoter interface that return a quote
type Quoter interface {
	Random() string
}

// POWer is a proof of work interface
type POWer interface {
	Validate(domain.Block) (bool, error)
	GetTask() *big.Int
}

// Chain is a blockchain interface
type Chain interface {
	Add(b domain.Block) bool
	All() []domain.Block
	Latest() domain.Block
}

// Client is a web service client
type Client struct {
	*typhon.Router
	pow    POWer
	chain  Chain
	quoter Quoter
}

// NewRoutes creates new service client
func NewRoutes(r *typhon.Router, pow POWer, chain Chain, q Quoter) *Client {
	return &Client{
		Router: r,
		pow:    pow,
		chain:  chain,
		quoter: q,
	}
}

// WithRoutes adds routes to the server
func (c *Client) WithRoutes() {
	c.Router.GET("/task", c.getTask)
	c.Router.GET("/chain", c.all)
	c.Router.POST("/verify", c.verify)
}

// getTask return a struct with task to resolve and block with only data
func (c Client) getTask(req typhon.Request) typhon.Response {
	block := domain.Block{}
	block.GenerateBody()
	block.PrevHash = c.chain.Latest().Hash

	response := struct {
		Task  *big.Int     `json:"task"`
		Block domain.Block `json:"block"`
	}{
		Task:  c.pow.GetTask(),
		Block: block,
	}
	return req.Response(response)
}

// verify is checking if hash of the block is calculated in the right way
func (c Client) verify(req typhon.Request) typhon.Response {
	var body domain.Block

	err := req.Decode(&body)
	if err != nil {
		return req.Response(err)
	}

	valid, err := c.pow.Validate(body)
	if err != nil {
		return req.ResponseWithCode(err, http.StatusInternalServerError)
	}

	if !valid {
		return req.ResponseWithCode(http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	if !c.chain.Add(body) {
		return req.ResponseWithCode(http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	response := struct {
		Quote string `json:"quote"`
	}{
		Quote: c.quoter.Random(),
	}

	return req.Response(response)
}

// all returns all blocks in the chain
// TODO: add pagination
func (c Client) all(req typhon.Request) typhon.Response {
	return req.Response(c.chain.All())
}
