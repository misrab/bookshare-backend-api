package handlers

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/coopernurse/gorp"

	"github.com/misrab/goapi/models"
)


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