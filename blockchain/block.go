package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/woohaen88/db"
	"github.com/woohaen88/utils"
)

const difficulty int = 2 // 0이 2개로 시작하는 hash
type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prev_hash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) generateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash + fmt.Sprint(b.Height)))
	b.Hash = fmt.Sprintf("%x", hash)
}

var ErrNotFound = errors.New("block not found!!!!")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	block.generateHash()
	block.persist()
	return &block

}

// POW 자격증명
// n개의 0으로 시작하는 hash를 찾음
