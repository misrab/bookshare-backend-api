package handlers


import (
	"time"
	
	"net/http"
	"encoding/json"

	"github.com/gorilla/schema"
)

var (
	// Recommended as package variable 
	// on http://www.gorillatoolkit.org/pkg/schema
	decoder = schema.NewDecoder()
	// HTTP header defaults
	HEADER_DEFAULTS = map[string]string {
		"Access-Control-Allow-Origin": "*",
		"Content-Type": "application/json",
	}
)


func SetHeaders(res http.ResponseWriter, code int) {
  	for k, v := range HEADER_DEFAULTS {
  		res.Header().Set(k, v)
  	}
  	res.Header().Set("Status", http.StatusText(code))
		res.Header().Set("Date", time.Now().String())	
}


func Respond(i interface{}, err error, res http.ResponseWriter) {
	// res.Write([]byte("fds"))
	// return

	if err != nil {
		SetHeaders(res, 400)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err2 := json.Marshal(i)
	if err2 != nil {
		SetHeaders(res, 400)
		http.Error(res, err2.Error(), http.StatusInternalServerError)
		return
	}

	SetHeaders(res, 200)
	res.Write(js)
}