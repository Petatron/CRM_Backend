package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/", fileServer)

	fmt.Println("Server is starting on port 3000. You can access it on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
