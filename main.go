package main

import (
	"fmt"
	"log"
	"net/http"


	"github.com/hello-world-twirp/pkg/paperclipserver"
	"github.com/hello-world-twirp/pkg/rpc"

)

func main() {
	fmt.Printf("Starting server on :6666")

	server := paperclipserver.NewServer()
	twirpHandler := rpc.NewUniversalPaperclipsServer(server, nil)

	log.Fatal(http.listenAndServe(":6666", twirpHandler))
}