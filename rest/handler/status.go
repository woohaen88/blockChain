package handler

import (
	"encoding/json"
	"net/http"

	"github.com/woohaen88/blockchain"
)

func Status(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blockchain())
}
