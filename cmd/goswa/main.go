package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	goswa "adnotanumber.com/goswa/pkg"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	var version bool
	flag.BoolVar(&version, "version", false, "display version and exit")
	flag.Parse()

	if version {
		fmt.Printf("version: %s\n", goswa.GetVersion())
		return nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	logger := goswa.NewAppLogger()
	server := goswa.NewServer(logger)
	logger.Info(fmt.Sprintf("Serving at %s:8080...", hostname))
	if err := http.ListenAndServe(":8080", server); err != nil {
		return err
	}
	return nil

}
