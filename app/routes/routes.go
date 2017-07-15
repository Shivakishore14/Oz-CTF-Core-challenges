package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/Shivakishore14/OzCTF-Challenge/app/console"
	"github.com/Shivakishore14/OzCTF-Challenge/app/controller"
	"github.com/gorilla/mux"
)

//LoadRoutes :for loading routing
func LoadRoutes() {
	address := ":8000"

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	challenge1 := r.PathPrefix("/challenge1").Subrouter()
	challenge1.HandleFunc("/{file}", controller.Challenge1)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	console.PrintSuccess("Listening on " + address)
	log.Fatal(srv.ListenAndServe())
}
