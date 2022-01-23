package main

import (
    "context"
    "net/http"
	"fmt"

   // pb "github.com/twitchtv/twirp-example/rpc/helloworld"
   "github.com/jbabinec1/hello-world-twirp/rpc"
)



// Run the implementation in a local server
func main() {
	fmt.Println("Client Example for Universal Paperclip Service")

	client := rpc.NewUniversalPaperclipsProtobufClient("http://localhost:6666", &http.Client{})
	_, err := client.IncrementPaperclips(context.Background(), &rpc.Size{Paperclips: 5})
	if err != nil {
		fmt.Println(err.Error())
	}
	paperclips, err := client.GetPaperclips(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Paperclips: ", paperclips.String())

	dread, err := client.CalculateUniverseLifespan(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Dread: ", dread.String())
}