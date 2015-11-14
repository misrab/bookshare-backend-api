package handlers

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	// "github.com/gorilla/mux"

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


func GetReadingsAutocomplete(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.Reading

	// vars := mux.Vars(req)
	// title := vars["title"]

	title := req.FormValue("title")


	query := "select * from readings where title ILIKE '"+title+"%' order by updated_at limit 10"
	_, err := dbmap.Select(&items, query)
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
	// item.Email = req.FormValue("email")
	// item.SetPassword(req.FormValue("password"))

	// save Reading
	err := dbmap.Insert(item)
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
	DeleteItem("readings", res, req, dbmap)
}
