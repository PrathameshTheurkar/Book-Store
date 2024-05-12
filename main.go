package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PrathameshTheurkar/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error while loading godotenv")
		panic(err)
	}

	fmt.Println("Hello from the main page")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
