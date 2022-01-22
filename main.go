package main

import (
	"fmt"
	"log"
	"net/http"

 
	"github.com/jbabinec1/hello-world-twirp/paperclipserver"
	"github.com/jbabinec1/hello-world-twirp/rpc"

)

func main() {
	fmt.Printf("Starting server on :6666")

	server := paperclipserver.NewServer()
	twirpHandler := rpc.NewUniversalPaperclipsServer(server, nil)

	log.Fatal(http.ListenAndServe(":6666", twirpHandler))
}