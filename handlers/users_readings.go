package handlers

import (
	// "fmt"
	"errors"
	"strconv"

	"net/http"
	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	"github.com/gorilla/mux"

	"github.com/misrab/bookshare-backend-api/models"
)

/*
    Curl test example

    curl -u UserReadingname:password localhost:8080/api/v1/UserReadings --data "email=aaa@aaa.com&password=123"

    curl -u UserReadingname:password localhost:8080/api/v1/UserReadings/1

    curl -u UserReadingname:password localhost:8080/api/v1/UserReadings/1 --request PATCH --data "email=bbb@bbb.com"

    curl -u UserReadingname:password localhost:8080/api/v1/UserReadings

    curl -u UserReadingname:password localhost:8080/api/v1/UserReadings/1 --request DELETE

    // get again to check
    curl -u UserReadingname:password localhost:8080/api/v1/UserReadings
*/





// func GetUserReadings(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
// 	var items []models.UserReading
// 	query := "select * from users_Readings order by updated_at"
// 	_, err := dbmap.Select(&items, query)
// 	if err != nil { 
// 		Respond(nil, err, res)
// 		return
// 	}

// 	Respond(items, err, res)
// }


// get the actual readings
func GetUserReadings(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var readings []models.Reading

	vars := mux.Vars(req)
	userId := vars["id"]

	_, err := dbmap.Select(&readings, "select * from readings where id in (select reading_id as id from users_readings where user_id = $1) order by updated_at desc", userId)
	
	Respond(readings, err, res)
}





// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostUserReading(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.UserReading)


	// get the user from auth
	err, user := GetUserFromAuth(req, dbmap)
	if err != nil { 
		Respond(nil, err, res)
		return
	}
	item.UserId = user.Id


	// TOOD if no reading id and article, create new article reading

	wrapper := struct{
		ReadingId int64 `json:"reading_id"`
	}{}
	err = json.NewDecoder(req.Body).Decode(&wrapper)
	if err != nil { 
		Respond(item, err, res)
		return
	}
	item.ReadingId = wrapper.ReadingId

	// item.ReadingId, err = Atoi64(req.Body["reading_id"])
	// if err != nil { 
	// 	Respond(item, err, res)
	// 	return
	// }

	// hmmm
	if (item.UserId == 0 || item.ReadingId == 0) {
		Respond(item, errors.New("Please provide a user_id and reading_id"), res)
		return
	}


	// save UserReading
	err = dbmap.Insert(item)
	Respond(item, err, res)
}


// func PatchUserReading(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
// 	var item models.UserReading
// 	modelname := "users_Readings"

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
	var err error
	for i, v := range items {
		result[i], err = Atoi64(v)
		if err != nil { return nil, err }
	}

	return result, nil
}

// func DeleteUserReading(modelname string, res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
// 	vars := mux.Vars(req)
// 	id := vars["id"]
// 	query := "delete from " + modelname + " where id=" + id
// 	_, err := dbmap.Exec(query)

// 	Respond(nil, err, res)
// }

// user from auth
func DeleteUserReading(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	// TODO delete actual reading if it's an article
	


	println("deleting user readgns")

	err, user := GetUserFromAuth(req, dbmap)
	if err != nil {
		Respond(nil, err, res)
		return
	}

	// fmt.Printf("%v\n", user)
	user_id := strconv.FormatInt(user.Id, 10)
	vars := mux.Vars(req)
	reading_id := vars["id"]

	query := "delete from users_readings where user_id=" + user_id + " and reading_id=" + reading_id
	_, err = dbmap.Exec(query)

	Respond(nil, err, res)
}
