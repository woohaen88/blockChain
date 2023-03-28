package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/woohaen88/db"
	"github.com/woohaen88/utils"
)

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prev_hash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
	Timestamp  int    `json:"timestamp"`
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

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("Target:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			return
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: Blockchain().difficulty(),
		Nonce:      0,
	}
	block.mine()
	block.persist()
	return &block

}

// POW 자격증명
// n개의 0으로 시작하는 hash를 찾음
