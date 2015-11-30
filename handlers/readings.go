package handlers

import (
	"net/http"
	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	// "github.com/gorilla/mux"
	"github.com/misrab/goutils"


	"github.com/misrab/bookshare-backend-api/models"
)

/*
    Curl test example

    curl -u Readingname:password localhost:8080/api/v1/Readings --data "email=aaa@aaa.com&password=123"

    curl -u Readingname:password localhost:8080/api/v1/Readings/1

    curl -u Readingname:password localhost:8080/api/v1/Readings/1 --request PATCH --data "email=bbb@bbb.com"

    curl -u Readingname:password localhost:8080/api/v1/Readings

    curl -u Readingname:password localhost:8080/api/v1/Readings/1 --request DELETE

    // get again to check
    curl -u Readingname:password localhost:8080/api/v1/Readings
*/

// example curl:
// 	curl -H "Content-Type: application/json" -X POST -d '{"url":"https://godoc.org/golang.org/x/net/html"}'  localhost:8000/api/v1/link_preview?token=lalala
// using post to pass url, which isn't pretty as a query parameter
func PostLinkPreview(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	// url := req.FormValue("url")
	item := struct {
		Url string `json:"url"`
	}{}

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil {
		Respond(item, err, res)
		return
	}
	url := item.Url


	// get the url
	preview, err := utils.ParseLink(url)
	
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(preview, err, res)

}


func GetReadingsAutocomplete(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.Reading

	// vars := mux.Vars(req)
	// title := vars["title"]

	title := req.FormValue("title") + "%"


	// query := "select * from readings where title ILIKE '"+title+"%' order by updated_at limit 10"
	_, err := dbmap.Select(&items, "select * from readings where title ILIKE $1 order by updated_at limit 10", title)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}



func GetReadings(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.Reading
	query := "select * from readings order by updated_at"
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}


func GetReading(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.Reading
	result, err := getById(item, req, dbmap)
	if err != nil { 
		Respond(nil, err, res)
		return
	}
	Respond(result, err, res)
}


// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostReading(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.Reading)

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil { 
		Respond(item, err, res)
		return
	}

	// save reading
	err = dbmap.Insert(item)
	Respond(item, err, res)
}


func PatchReading(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.Reading
	modelname := "readings"

	id, _ := getId(req)
	query := "select * from " + modelname + " where id=$1"
	err := dbmap.SelectOne(&item, query, id)
	if err != nil {
		Respond(nil, err, res)
		return
	}

	err = parseFormValues(&item, req)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	_, err2 := dbmap.Update(&item)
	Respond(item, err2, res)
	
}

func DeleteReading(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	// TODO clear associations
	
	DeleteItem("readings", res, req, dbmap)
}
