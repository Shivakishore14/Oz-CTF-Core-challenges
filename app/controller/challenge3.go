package controller

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/Shivakishore14/OzCTF-Challenge/app/model"
)

func Challenge3(w http.ResponseWriter, r *http.Request) {
	mike := model.C3User{Name: "mike", Hobbies: "Gaming", Secret: "I fear cats"}
	dan := model.C3User{Name: "dam", Hobbies: "photography", Secret: "no secrets, am an open book :)"}
	harry := model.C3User{Name: "harry", Hobbies: "Reading books", Secret: "Am really a wizard B)"}
	kate := model.C3User{Name: "kate", Hobbies: "Gaming", Secret: "OZCTF{flag3}"}
	t := template.New("example")
	t, _ = t.Parse(`
		<html>
			<center>
				<h1>{{.Name}}<h1><br>
				<p> Hobbies : {{.Hobbies}} </p> <br>
				<div style="display:none">secret : {{.Secret}}</div>
			</center>
		</html>
		`)
	switch r.FormValue("profile") {
	case "mike":
		t.Execute(w, mike)
		break
	case "dan":
		t.Execute(w, dan)
		break
	case "harry":
		t.Execute(w, harry)
		break
	case "kate":
		t.Execute(w, kate)
	default:
		htmlFile, err := ioutil.ReadFile("./challenge3/welcome.html")
		if err != nil {
			fmt.Print(err)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(htmlFile)
	}
}
