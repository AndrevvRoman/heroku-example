package main

import (
	"fmt"
	"heroku-example/pkg/transport"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

const serverURL = ":8080"

func main() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{
		"url": serverURL,
	}).Info("starting the server")
	r := transport.Router()
	fmt.Println(http.ListenAndServe(serverURL, r))
}
