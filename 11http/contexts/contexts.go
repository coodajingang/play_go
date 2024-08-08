package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		context := r.Context()
		fmt.Println("hello start")
		defer fmt.Println("hello end")
		select {
		case <-context.Done():
			err := context.Err()
			fmt.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		case <-time.After(10 * time.Second):
			fmt.Fprintf(w, "hello end\n")
		}
	})

	http.ListenAndServe(":8080", nil)
}
