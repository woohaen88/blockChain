package main

import (
	"github.com/woohaen88/blockchain"
	"github.com/woohaen88/rest"
)

func main() {
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Thrid")
	rest.Start(3000)
}
