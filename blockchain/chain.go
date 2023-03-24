package blockchain

import (
	"sync"
)



type blockchain struct {
	Newesthash string `json:"newest_hash"`
	Height int `json:"height"`
}

var b *blockchain
var once sync.Once




func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.Newesthash, b.Height)
	b.Newesthash = block.Hash
	b.Height = block.Height
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis")
		})
	}
	return b
}