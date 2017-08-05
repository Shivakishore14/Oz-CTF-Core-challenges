package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Challenge2(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("flag")
	data := []byte("OZCTF{C00k1e5_Ar3_N0t_S0_S3cUr3}")

	b64flag := base64.StdEncoding.EncodeToString(data)
	w.Header().Set("Set-Cookie", "flag="+b64flag)

	htmlFile, err := ioutil.ReadFile("./challenge2/welcome.html")
	if err != nil {
		fmt.Print(err)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write(htmlFile)
}
