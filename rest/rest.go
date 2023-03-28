package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/woohaen88/rest/handler"
)

var Port string

func Start(aPort int) {
	Port = fmt.Sprintf(":%d", aPort)
	route := mux.NewRouter()
	route.Use(ApplicationJsonMiddleWare)
	route.HandleFunc("/", handler.Home)
	route.HandleFunc("/block", handler.Block)
	route.HandleFunc("/block/{hash:[a-f0-9]+}", handler.GetBlockHeight)
	route.HandleFunc("/status", handler.Status)
	fmt.Printf("Listen Server http://localhost%s\n", Port)
	http.ListenAndServe(":3000", route)

}
