package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
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
			fmt.Printf("NewestHash: %s\nHeight:%d\n", b.Newesthash, b.Height)
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			}else{
				fmt.Println("Restoring...")
				b.restore(checkpoint)
			}
			
		})
	}
	fmt.Printf("NewestHash: %s\nHeight:%d\n", b.Newesthash, b.Height)
	return b
}

func (b *blockchain) restore(data []byte){
	decoder := gob.NewDecoder(bytes.NewReader(data))
	utils.HandleErr(decoder.Decode(b))
}