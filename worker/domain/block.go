package domain

import "math/big"

type Block struct {
	Hash     []byte `json:"hash"`
	Data     []byte `json:"data"`
	PrevHash []byte `json:"prev_hash"`
	Nonce    int    `json:"nonce"`
}

type TaskResponse struct {
	Task  *big.Int `json:"task"`
	Block *Block   `json:"block"`
}

type VerifyResponse struct {
	Quote string `json:"quote"`
}
