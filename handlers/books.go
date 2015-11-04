package handlers

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	"github.com/misrab/bookshare-backend-api/models"
)

/*
    Curl test example

    curl -u Bookname:password localhost:8080/api/v1/Books --data "email=aaa@aaa.com&password=123"

    curl -u Bookname:password localhost:8080/api/v1/Books/1

    curl -u Bookname:password localhost:8080/api/v1/Books/1 --request PATCH --data "email=bbb@bbb.com"

    curl -u Bookname:password localhost:8080/api/v1/Books

    curl -u Bookname:password localhost:8080/api/v1/Books/1 --request DELETE

    // get again to check
    curl -u Bookname:password localhost:8080/api/v1/Books
*/


func GetBooksAutocomplete(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.Book
	title := req.FormValue("title")
	query := "select * from Books where title like '%"+title+"%' order by updated_at limit 10"
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}



func GetBooks(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.Book
	query := "select * from books order by updated_at"
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}


func GetBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.Book
	result, err := getById(item, req, dbmap)
	if err != nil { 
		Respond(nil, err, res)
		return
	}
	Respond(result, err, res)
}


// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.Book)
	// item.Email = req.FormValue("email")
	// item.SetPassword(req.FormValue("password"))

	// save Book
	err := dbmap.Insert(item)
	Respond(item, err, res)
}


func PatchBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.Book
	modelname := "books"

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

func DeleteBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	DeleteItem("books", res, req, dbmap)
}
