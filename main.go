package main

import (
	"crypto/sha256"
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

func main() {
	chain := GetBlockchain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.AllBlocks() {
		fmt.Println("=========================")
		fmt.Println("Data: ", block.Data)
		fmt.Println("Hash: ", block.Hash)
		fmt.Println("PrevHash: ", block.PrevHash)
		fmt.Println("Height: ", block.Height)
	}
}
