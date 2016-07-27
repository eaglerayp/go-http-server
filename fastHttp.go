package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	ln, err := net.Listen("tcp4", ":8888")
	if err != nil {
		log.Fatalf("error in net.Listen: %s", err)
	}
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Printf("")
	}

	// Start the server with default settings.
	// Create Server instance for adjusting server settings.
	//
	// Serve returns on ln.Close() or error, so usually it blocks forever.
	if err := fasthttp.Serve(ln, requestHandler); err != nil {
		log.Fatalf("error in Serve: %s", err)
	}
}
