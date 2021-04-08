package transport

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//Router bla bla lba
func Router() http.Handler {
	log.Info("creating router")
	r := mux.NewRouter()
	// s := r.PathPrefix("api/v1").Subrouter()
	r.HandleFunc("/hello-world", helloWorld)
	return logMiddleware(r)
}

func logMiddleware(h http.Handler) http.Handler {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
	return h
}

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Hello user")
}
