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





// func GetUserBooks(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
// 	var items []models.UserBook
// 	query := "select * from users_books order by updated_at"
// 	_, err := dbmap.Select(&items, query)
// 	if err != nil { 
// 		Respond(nil, err, res)
// 		return
// 	}

// 	Respond(items, err, res)
// }


// func GetUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
// 	var item models.UserBook
// 	result, err := getById(item, req, dbmap)
// 	if err != nil { 
// 		Respond(nil, err, res)
// 		return
// 	}
// 	Respond(result, err, res)
// }



// func associateUserBook(userId, bookId int64) error {

// }

// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.UserBook)
	var err error

	item.UserId, err = Atoi64(req.FormValue("user_id"))
	if err != nil { 
		Respond(item, err, res)
		return
	}
	item.BookId, err = Atoi64(req.FormValue("book_id"))
	if err != nil { 
		Respond(item, err, res)
		return
	}

	// item.SetPassword(req.FormValue("password"))
	if (item.UserId == 0 || item.BookId == 0) {
		Respond(item, errors.New("Please provide a user_id and book_id"), res)
		return
	}

	// TODO verify they're valid / not already associated


	// save UserBook
	err = dbmap.Insert(item)
	Respond(item, err, res)
}


// func PatchUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
// 	var item models.UserBook
// 	modelname := "users_books"

// 	id, _ := getId(req)
// 	query := "select * from " + modelname + " where id=$1"
// 	err := dbmap.SelectOne(&item, query, id)
// 	if err != nil {
// 		Respond(nil, err, res)
// 		return
// 	}

// 	err = parseFormValues(&item, req)
// 	if err != nil { 
// 		Respond(nil, err, res)
// 		return
// 	}

// 	_, err2 := dbmap.Update(&item)
// 	Respond(item, err2, res)
	
// }

func Atoi64(s string) (i int64, err error) {
    i64, err := strconv.ParseInt(s, 10, 64)
    return int64(i64), err
}

func StringsToInts64(items ...string) ([]int64, error) {
	result := make([]int64, len(items))
	for _, v := range items {
		item.UserId, err = Atoi64(v)
	}
}

func DeleteItem(modelname string, res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	vars := mux.Vars(req)
	id := vars["id"]
	query := "delete from " + modelname + " where id=" + id
	_, err := dbmap.Exec(query)

	Respond(nil, err, res)
}

func DeleteUserBook(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	vars := mux.Vars(req)
	id := vars["user_id"]
	query := "delete from " + modelname + " where id=" + id
	_, err := dbmap.Exec(query)

	Respond(nil, err, res)
}
