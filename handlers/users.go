package handlers

import (
	"errors"
	// "fmt"

	"net/http"
	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	// "github.com/gorilla/mux"

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


/*
	Non-handlers
*/

type UserWrapper struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserAuthWrapper struct {
	Email string `json:"email"`
	Hash string `json:"Hash"`
}


// decodes user creds from auth header and tries to 
// find such user. If no valid match returns nil.
func GetUserFromAuth(req *http.Request, dbmap *gorp.DbMap) (error, *models.User) {
	var item models.User
	username, hash, httpCode := DecodeAuthHeader(req)

	if httpCode != http.StatusOK { return errors.New("Invalid auth"), nil }

	// get said user
	err := dbmap.SelectOne(&item, "select * from users where email=$1", username)
	// check hash
	if err != nil || item.Hash != hash { return err, nil }

	return nil, &item
}



/*
	Handlers
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


func GetCurrentUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	// var item models.User

	// // first decode header
	// username, _, httpCode := DecodeAuthHeader(req)
	// if httpCode != http.StatusOK {
	// 	// errors.New("Invalid authorization header")
	// 	Respond(nil, errors.New("Invalid authorization header"), res)
	// }

	// // get said user
	// err := dbmap.SelectOne(&item, "select * from users where email=$1", username)
	// if err != nil { 
	// 	// err
	// 	Respond(nil, err, res)
	// 	return
	// }

	// check password TODO
	err, user := GetUserFromAuth(req, dbmap)


	Respond(user, err, res)
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

func LoginUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	// get email and password
	item := new(models.User)
	// vars := mux.Vars(req)
	// email := vars["email"]
	// password := vars["password"]

	var userWrapper UserWrapper
	var err error
	
	err = json.NewDecoder(req.Body).Decode(&userWrapper)
	if err != nil {
		Respond(nil, err, res)
		return
	}

	// get the user
	err = dbmap.SelectOne(&item, "select * from users where email=$1", userWrapper.Email)
	if err != nil {
		Respond(nil, err, res)
		return
	}

	// check the password
	err = item.ComparePassword(userWrapper.Password)
	if err != nil {
		Respond(nil, err, res)
		return
	}

	Respond(item, nil, res)
}




// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostUser(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.User)


	var userWrapper UserWrapper
	var err error
	
	err = json.NewDecoder(req.Body).Decode(&userWrapper)
	if err != nil {
		Respond(nil, err, res)
		return
	}


	email := userWrapper.Email
	password := userWrapper.Password

	// set new object for insert
	item.Email = email
	item.SetPassword(password)


	// save user
	err = dbmap.Insert(item)
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
