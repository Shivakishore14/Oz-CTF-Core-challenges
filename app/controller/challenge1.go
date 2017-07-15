package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Challenge1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		file := vars["file"]
		fmt.Print(file)
		if file == "" || file == "admin.html" {
			file = "login.html"
		}
		htmlFile, err := ioutil.ReadFile("./challenge1/" + file)
		if err != nil {
			fmt.Print(err)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(htmlFile)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "Adm1n" && password == "p@ssw0rd_@dM!N" {
			htmlFile, err := ioutil.ReadFile("./challenge1/admin.html")
			if err != nil {
				fmt.Print(err)
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Write(htmlFile)
		} else {
			fmt.Fprint(w, "provide correct credentials")
		}
	}
}
