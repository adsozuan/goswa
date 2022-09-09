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
	logger := goswa.NewAppLogger()
	server := goswa.NewServer(logger)
	logger.Info("Serving at :8080...")
	if err := http.ListenAndServe(":8080", server); err != nil {
		return err
	}
	return nil

}
