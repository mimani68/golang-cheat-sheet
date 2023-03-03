package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// https://www.sohamkamani.com/golang/context-cancellation-and-values/

//
// Scenario:
// -------------------------------------
//    curl http://localhost:8000
//    after 2 seconds perform CANSELATION by using "ctl+c"
//
func main() {
	// Total pocessing time for completing single Request
	const operationDelay = 5 * time.Second
	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Fprint(os.Stdout, "[INFO] Start processing new request\n")
		// We use `select` to execute a piece of code depending on which
		// channel receives a message first
		select {
		case <-time.After(operationDelay):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("{\"status\":\"request processed\"}"))
			fmt.Fprint(os.Stdout, "[INFO] Responce sended \n")
		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "[ERROR] Request cancelled\n")
		}
		fmt.Fprint(os.Stdout, "[INFO] FINISH \n")
	}))
}
