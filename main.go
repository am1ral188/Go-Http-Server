package main

import (
	"awesomeProject/src/cfg"
	"awesomeProject/src/router"
	"awesomeProject/src/tools"
	"fmt"
	"net/http"
)

func main() {
	//fixme :add sub path handler
	//fixme : replace all fmt print with log.New
	mux := tools.Mux

	server := &http.Server{Addr: cfg.ListenOn, Handler: mux}

	fs := http.FileServer(http.Dir("./src/view/assets"))

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	router.RouterInit()

	listenError := server.ListenAndServe()
	if listenError != nil {
		fmt.Printf("faild to listen on port: %t err:%t\n", cfg.ListenOn, listenError)
	}

}
