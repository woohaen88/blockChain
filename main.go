package main

import (
	"github.com/woohaen88/db"
	"github.com/woohaen88/rest"
)

func main() {
	defer db.Close()	
	rest.Start(3000)
}
