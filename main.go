package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func roothandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "hi hello",
	})
}

func main() {
	fmt.Println("server is running")
	log.Fatal(http.ListenAndServe("localhost:9000", http.HandlerFunc(roothandler)))

}
