package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/stream")
	if err != nil {
		fmt.Println("Error connecting to server")
	}

	defer resp.Body.Close() // release resources when func is done
	// kind of an async free

	buf := make([]byte, 1024) // buffer to temp store messages from server
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF {
			fmt.Println("End of stream")
			break
		}
		if err != nil {
			fmt.Println("Error in stream", err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
