package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/readyz", ready)
	log.Printf("Starting")
	err := http.ListenAndServe(":8080", nil)
	log.Printf("Terminating [%s]", err.Error())
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(100) == 42 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Hello %s\n", r.Host)
}

func readyz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}
