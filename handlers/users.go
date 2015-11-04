package handlers

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	"github.com/misrab/bookshare-backend-api/models"
)

/*
    Curl test example

    curl -u username:password localhost:8080/api/v1/users --data "email=aaa@aaa.com&password=123"

    curl -u username:password localhost:8080/api/v1/users/1

    curl -u username:password localhost:8080/api/v1/users/1 --request PATCH --data "email=bbb@bbb.com"

    curl -u username:password localhost:8080/api/v1/users

    curl -u username:password localhost:8080/api/v1/users/1 --request DELETE

    // get again to check
    curl -u username:password localhost:8080/api/v1/users
*/



func GetUsers(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.User
	query := "select * from users order by updatedAt"
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}


func GetUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.User
	result, err := getById(item, req, dbmap)
	if err != nil { 
		Respond(nil, err, res)
		return
	}
	Respond(result, err, res)
}


// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.User)
	item.Email = req.FormValue("email")
	item.SetPassword(req.FormValue("password"))

	// save user
	err := dbmap.Insert(item)
	Respond(item, err, res)
}


func PatchUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.User
	modelname := "users"

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

func DeleteUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	DeleteItem("users", res, req, dbmap)
}
