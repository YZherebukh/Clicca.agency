package domain

import "math/rand"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Block is a block struct
type Block struct {
	Hash     []byte `json:"hash"`
	Data     []byte `json:"data"`
	PrevHash []byte `json:"prev_hash"`
	Nonce    int    `json:"nonce"`
}

// GenerateBody randomly generates block body for simplicity
func (b *Block) GenerateBody() {
	body := make([]rune, rand.Intn(len(letterRunes)))
	for i := range body {
		body[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	b.Data = []byte(string(body))
}
