package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/RemiPail/moncha/pkg/db"
	"github.com/RemiPail/moncha/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/shopItem", h.GetAllShopItems).Methods(http.MethodGet)
	router.HandleFunc("/shopItem", h.AddShopItem).Methods(http.MethodPost)

	port := os.Getenv("PORT")
	log.Println("API is running on port " + port + "!")
	http.ListenAndServe(":"+port, router)
}
