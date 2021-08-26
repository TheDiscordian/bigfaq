package main

import (
	"net/http"
	"log"
)

// TODO Also search questions when querying

func main() {
	connectDB()
	initHandlers()
	log.Fatal(http.ListenAndServe(":8585", nil))
}
