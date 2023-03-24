package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prev_hash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
	Blocks []*Block
}

var b *blockchain
var once sync.Once

func getLashHash() string {
	totalLenth := len(b.Blocks)
	if totalLenth > 0 {
		return b.Blocks[totalLenth-1].Hash
	}
	return ""
}

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash + strconv.Itoa(b.Height)))
	b.Hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLashHash(), len(b.Blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.Blocks = append(b.Blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return GetBlockchain().Blocks
}

var NotFoundError = errors.New("Not Found Error")

func (b *blockchain) Height(height string) (*Block, error) {
	heightInt, _ := strconv.Atoi(height)	

	for _, block := range b.Blocks {
		if heightInt == block.Height{
			return block, nil
		}
	}
	
	return nil, NotFoundError
}
