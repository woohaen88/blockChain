package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/woohaen88/blockchain"
)

type AddBlockBody struct {
	Data string `json:"data"`
}

func Block(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		json.NewDecoder(r.Body).Decode(&addBlockBody)
		blockchain.GetBlockchain().AddBlock(addBlockBody.Data)
		w.WriteHeader(http.StatusCreated)
	}
}

func GetBlockHeight(w http.ResponseWriter, r *http.Request) {
	// block 데이터 한개 가져오기
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	block, _ := blockchain.GetBlockchain().Height(vars["height"])

	json.NewEncoder(w).Encode(block)
}
