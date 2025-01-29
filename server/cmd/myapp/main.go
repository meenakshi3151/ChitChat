package main

import (
	"ChitChat/internal/route"
	"fmt"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"ChitChat/internal/config"
)
func main() {
	fmt.Println("Chit Chat Server!!")
	config.LoadEnv()
	config.Db()
	PORT := os.Getenv("PORT")
	fmt.Println("Server running on http://localhost" + PORT)
	router := mux.NewRouter()
	route.UserRoutes(router)
	log.Fatal(http.ListenAndServe(PORT, router))
}