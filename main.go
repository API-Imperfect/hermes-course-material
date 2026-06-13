package main

import (
	"log"
	"net/http"

	"sample/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Register one Register<Resource> per resource file in handlers/.
	handlers.RegisterUsers(mux)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
