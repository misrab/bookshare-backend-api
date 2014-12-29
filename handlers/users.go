package handlers

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/coopernurse/gorp"

	"github.com/misrab/goapi/models"
)

/*
	Users Handlers:
	- GET many
	- GET one
	- PATCH one
	- POST one
	- DELETE one
*/



func GetUsers(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.User
	query := "select * from users order by updated"
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}


func GetUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.User
	result, err := GetById(item, req, dbmap)
	if err != nil { 
		Respond(nil, err, res)
		return
	}
	Respond(result, err, res)
}



func PostUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.User)
	item.Email = req.FormValue("email")
	item.SetPassword(req.FormValue("password"))

	// save user
	err := dbmap.Insert(item)
	Respond(item, err, res)
}


func DeleteUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	DeleteItem("users", res, req, dbmap)
}
