package main

import (
	"log"
	"net/http"

	"github.com/chrisedrego/gitveaver/internal/handlers"
	"github.com/chrisedrego/gitveaver/utils"
)

type error interface {
	Error() string
}

func main() {
	utils.FlagCheck()
	http.HandleFunc("/", handlers.RequestHandler)
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
