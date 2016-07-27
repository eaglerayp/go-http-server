package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", testHttp)

	http.ListenAndServe(":8888", nil)
}

func testHttp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the resource.
		fmt.Fprintf(w, `{"Code":0}`)
	case "POST":
		// Create a new record.
		fmt.Fprintf(w, "")
	default:
		// Give an error message.
	}
}
