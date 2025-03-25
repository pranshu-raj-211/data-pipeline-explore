package main

import (
	"fmt"
	"net/http"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Client connected")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "data: Message %d \n\n", i)
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}
	fmt.Println("End of stream, closing connection")
}

func main() {
	http.HandleFunc("/stream", streamHandler)
	fmt.Println("setting up server")
	http.ListenAndServe(":8080", nil)
}
