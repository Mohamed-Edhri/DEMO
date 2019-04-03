package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test app!!\n")
}

func main() {
	log.Print("test app server ready")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50051", nil)
}
