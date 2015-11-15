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
		"Access-Control-Allow-Headers": "Origin,Content-Type,Accept,Authorization",
		"Content-Type": "application/json",
		"Access-Control-Allow-Methods": "GET,PATCH,PUT,POST,DELETE",
	}
)

// res.Header().Set("Access-Control-Allow-Origin", "*")


func SetHeaders(res http.ResponseWriter, code int) {
	// println("setting headers")

  	for k, v := range HEADER_DEFAULTS {
  		res.Header().Set(k, v)
  	}

  	// res.Header().Set("Access-Control-Allow-Origin", "*")
  	res.Header().Set("Status", http.StatusText(code))
	res.Header().Set("Date", time.Now().String())	
}


func Respond(i interface{}, err error, res http.ResponseWriter) {
	// empty result
	if (err != nil && err.Error() == "sql: no rows in result set") {
		SetHeaders(res, http.StatusOK)
		// res.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(res).Encode(i); err != nil {
	        json.NewEncoder(res).Encode(err)
	    }
		return
	}

	if err != nil {
		SetHeaders(res, http.StatusBadRequest)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// js, err2 := json.Marshal(i)
	// if err2 != nil {
	// 	SetHeaders(res, http.StatusBadRequest)
	// 	http.Error(res, err2.Error(), http.StatusBadRequest)
	// 	return
	// }

	// println(js)


	SetHeaders(res, http.StatusOK)
	if err := json.NewEncoder(res).Encode(i); err != nil {
        json.NewEncoder(res).Encode(err)
    }
	// res.Write(js)
}