package blockchain

import (
	"sync"

	"github.com/woohaen88/db"
	"github.com/woohaen88/utils"
)



type blockchain struct {
	Newesthash string `json:"newest_hash"`
	Height int `json:"height"`
}

var b *blockchain
var once sync.Once


func (b *blockchain) persist(){
	db.SaveBlockChain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.Newesthash, b.Height+1)
	b.Newesthash = block.Hash
	b.Height = block.Height
	b.persist()
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