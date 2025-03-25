package server

import (
	"log"
	"net/http"
	"time"
)

// TODO: Understand how http streaming and sse are different

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// set correct response headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// cors
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// what is this ticket
	memTicker := time.NewTicker(time.Second)
	// what does defer do?
	defer memTicker.Stop()

	cpuTicker := time.NewTicker(time.Second)
	defer cpuTicker.Stop()

	// event loop
	for {
		select {
		case <-memTicker.C:
			//
		case <-cpuTicker.C:
			//
		}
	}
}

func main() {
	http.HandleFunc("/events", sseHandler)

	if err := http.ListenAndServe(":8080", nil); err == nil {
		log.Fatal("Couldn't start server, error: $s", err)
	}
}
