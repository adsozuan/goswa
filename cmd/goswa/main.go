package main

import (
	"log"
	"net/http"

	"adnotanumber.com/goswa"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server := goswa.NewServer()
	if err := http.ListenAndServe(":8080", server); err != nil {
		return err
	}
	return nil

}
