package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

func main() {
	ln, err := reuseport.Listen("tcp4", ":8888")
	if err != nil {
		log.Fatalf("error in reuseport listener: %s", err)
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
