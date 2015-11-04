package handlers

import (
	"errors"
	"strconv"

	"net/http"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	"github.com/misrab/bookshare-backend-api/models"
)

/*
    Curl test example

    curl -u UserBookname:password localhost:8080/api/v1/UserBooks --data "email=aaa@aaa.com&password=123"

    curl -u UserBookname:password localhost:8080/api/v1/UserBooks/1

    curl -u UserBookname:password localhost:8080/api/v1/UserBooks/1 --request PATCH --data "email=bbb@bbb.com"

    curl -u UserBookname:password localhost:8080/api/v1/UserBooks

    curl -u UserBookname:password localhost:8080/api/v1/UserBooks/1 --request DELETE

    // get again to check
    curl -u UserBookname:password localhost:8080/api/v1/UserBooks
*/


func GetUserBooksAutocomplete(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.UserBook
	title := req.FormValue("title")
	query := "select * from users_books where title like '%"+title+"%' order by updated_at limit 10"
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}



func GetUserBooks(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.UserBook
	query := "select * from users_books order by updated_at"
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}


func GetUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.UserBook
	result, err := getById(item, req, dbmap)
	if err != nil { 
		Respond(nil, err, res)
		return
	}
	Respond(result, err, res)
}


// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.UserBook)
	var err error
	item.UserId, err = strconv.Atoi(req.FormValue("user_id"))
	if err != nil { 
		Respond(item, err, res)
		return
	}
	item.BookId, err = strconv.Atoi(req.FormValue("book_id"))
	if err != nil { 
		Respond(item, err, res)
		return
	}

	// item.SetPassword(req.FormValue("password"))
	if (item.UserId == 0 || item.UserId == 0) {
		Respond(item, errors.New("Please provide a user_id and book_id"), res)
		return
	}

	// TODO verify they're valid / not already associated

	// save UserBook
	err := dbmap.Insert(item)
	Respond(item, err, res)
}


func PatchUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.UserBook
	modelname := "users_books"

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

func DeleteUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	DeleteItem("users_books", res, req, dbmap)
}
