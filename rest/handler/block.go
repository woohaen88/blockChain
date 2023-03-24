package handler

import (
	"encoding/json"
	"fmt"
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
		//json.NewEncoder(w).Encode(blockchain.GetBlockchain().AllBlocks())
		return
	case "POST":
		return 
		// var addBlockBody AddBlockBody
		// json.NewDecoder(r.Body).Decode(&addBlockBody)
		// blockchain.GetBlockchain().AddBlock(addBlockBody.Data)
		// w.WriteHeader(http.StatusCreated)
	}
}

type Error struct {
	Message string `json:"message"`
}

func GetBlockHeight(w http.ResponseWriter, r *http.Request) {
	// block 데이터 한개 가져오기
	vars := mux.Vars(r)
	hash := vars["hash"]
	w.WriteHeader(http.StatusOK)
	block, err := blockchain.FindBlock(hash)
	encode := json.NewEncoder(w)
	if err == blockchain.ErrNotFound {
		encode.Encode(Error{Message: fmt.Sprint(err)})
	}else{
		encode.Encode(block)
	}

	
}
