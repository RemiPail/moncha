package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	randomHandler := func(w http.ResponseWriter, r *http.Request) {
		var choices []string
		err := json.NewDecoder(r.Body).Decode(&choices)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pick := choices[rand.Intn(len(choices))]
		io.WriteString(w, pick)
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/random", randomHandler)
	log.Println("Listening for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
