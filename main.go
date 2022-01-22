package main

import (
	"fmt"
	"log"
	"net/http"

 
	"github.com/jbabinec1/hello-world-twirp/paperclipserver"
	//"github.com/jbabinec1/hello-world-twirp/pkg/rpc"
	"github.com/jbabinec1/hello-world-twirp/pkg"

)

func main() {
	fmt.Printf("Starting server on :6666")

	server := paperclipserver.NewServer()
	twirpHandler := rpc.NewUniversalPaperclipsServer(server, nil)

	log.Fatal(http.listenAndServe(":6666", twirpHandler))
}