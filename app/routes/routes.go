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
	r.PathPrefix("/challenge0/").Handler(http.StripPrefix("/challenge0/", http.FileServer(http.Dir("challenge0"))))
	r.PathPrefix("/challenge5/").Handler(http.StripPrefix("/challenge5/", http.FileServer(http.Dir("challenge5"))))
	r.PathPrefix("/challenge6/").Handler(http.StripPrefix("/challenge6/", http.FileServer(http.Dir("challenge6"))))
	r.PathPrefix("/challenge7/").Handler(http.StripPrefix("/challenge7/", http.FileServer(http.Dir("challenge7"))))

	challenge1 := r.PathPrefix("/challenge1").Subrouter()
	challenge1.HandleFunc("/{file}", controller.Challenge1)

	challenge2 := r.PathPrefix("/challenge2").Subrouter()
	challenge2.HandleFunc("/", controller.Challenge2)

	challenge3 := r.PathPrefix("/challenge3").Subrouter()
	challenge3.HandleFunc("/", controller.Challenge3)

	challenge4 := r.PathPrefix("/challenge4").Subrouter()
	challenge4.HandleFunc("/", controller.Challenge4)
	challenge4.HandleFunc("/set", controller.Challenge4SetScore)
	challenge4.HandleFunc("/getkey", controller.Challenge4GetFlag)

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	console.PrintSuccess("Listening on " + address)
	log.Fatal(srv.ListenAndServe())
}
