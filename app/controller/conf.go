package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shivakishore14/OzCTF-Challenge/app/console"
	"github.com/Shivakishore14/OzCTF-Challenge/app/model"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

var database = "coda"
var user = "root"
var password = "sken"

var db *gorm.DB
var store *sessions.CookieStore

func init() {
	var err error

	store = sessions.NewCookieStore([]byte("secret-key-to-be-changed"))
	if db, err = gorm.Open("mysql", user+":"+password+"@/"+database+"?charset=utf8&parseTime=True&loc=Local"); err != nil {
		console.PrintError("Error Connecting to database")
	}

	fmt.Println("DEBUG : Init Contollers Done")
}

func webresponse(msg string, err error, data interface{}, w http.ResponseWriter) (string, error) {
	obj := model.WebResponse{}
	obj.Message = msg
	obj.Error = err
	obj.Data = data

	var resErr error
	var resTxt string

	if jsonData, e := json.Marshal(obj); e != nil {
		resErr = e
	} else {
		resTxt = string(jsonData)
	}
	fmt.Fprint(w, resTxt)
	return resTxt, resErr
}
