package blockchain

import (
	"time"

	"github.com/woohaen88/utils"
)

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"tx_ins"`
	TxOuts    []*TxOut `json:"tx_outs"`
}

const (
	minerReward int = 50
)

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Owner  string
	Amount int
}

func (t *Tx) setId() {
	// Tx에서 Id를 해쉬
	t.Id = utils.Hash(t)
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := &Tx{
		Id:        "", // 나중에 거래내역을 Hash하면서 지정
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.setId()
	return tx
}
