package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PrathameshTheurkar/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error while Loading godotenv")
		panic(err)
	}

	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))

}
