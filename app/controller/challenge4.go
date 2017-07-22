package controller

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Challenge4(w http.ResponseWriter, r *http.Request) {

	htmlFile, err := ioutil.ReadFile("./challenge4/game.html")
	if err != nil {
		fmt.Print(err)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(htmlFile)

}

func Challenge4SetScore(w http.ResponseWriter, r *http.Request) {
	scoreVal := r.FormValue("score")
	score, e := strconv.Atoi(scoreVal)
	if e != nil {
		fmt.Fprint(w, "not valid")
		return
	}
	keyEncrypted := "YnZybnB5cGt4dA=="
	if score > 999999 {
		t := template.New("example")
		t, _ = t.Parse(`
			<html>
				<center>
					<h1>Great</h1><br>
					<i>Based</i> on your score i think you shoud <i>Playfair</i> <br> anyway have this [ ` + keyEncrypted + ` ]
					<form action="getkey" method="POST">
						<input type="hidden" name="key"/>
					</form>
				</center>
			</html>
			`)
		t.Execute(w, nil)
	} else {
		fmt.Fprint(w, "try harder")
	}
}
func Challenge4GetFlag(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if r.Method == "POST" {
		if key == "awsmozkeys" {
			fmt.Fprint(w, "OZCTF{flag4}")
		} else {
			fmt.Fprint(w, "wrong key")
		}
	} else {
		http.Redirect(w, r, "/challenge4/", 302)
	}

}
