package main

import (
	"log"
	"net/http"

	"github.com/collings22/knowledge-oppo/routes"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.RoutesHandler(r)

	r.Use(mux.CORSMethodMiddleware(r))

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))

}
