package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	randomHandler := func(w http.ResponseWriter, req *http.Request) {
		choices := [3]string{"Pile", "Face", "Dans ton cul !!!"}
		pick := choices[rand.Intn(len(choices))]

		io.WriteString(w, pick)
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/random", randomHandler)
	log.Println("Listening for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
