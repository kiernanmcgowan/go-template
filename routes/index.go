package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BindRoutes binds routes to the processes's http handler
func BindRoutes() {
	// set up our router
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/{domain}", company)

	// and bind the router to the http handler
	http.Handle("/", r)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Tracing request for %s", r.RequestURI)
	w.Write([]byte("Hello world!"))
}

func company(w http.ResponseWriter, r *http.Request) {
	log.Printf("Tracing request for %s", r.RequestURI)
	vars := mux.Vars(r)
	w.Write([]byte(fmt.Sprintf("You're trying to lookup %s", vars["domain"])))
}
