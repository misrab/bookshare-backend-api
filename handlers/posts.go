package handlers

import (
	"errors"
	"sync"
	// "fmt"

	"net/http"
	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

	// "github.com/gorilla/mux"

	"github.com/misrab/bookshare-backend-api/models"
)


const (
	POST_OFFSET_INCREMENT = "10" // mirrored on client
)

/*
    Curl test example

    curl -u Postname:password localhost:8080/api/v1/Posts --data "email=aaa@aaa.com&password=123"

    curl -u Postname:password localhost:8080/api/v1/Posts/1

    curl -u Postname:password localhost:8080/api/v1/Posts/1 --request PATCH --data "email=bbb@bbb.com"

    curl -u Postname:password localhost:8080/api/v1/Posts

    curl -u Postname:password localhost:8080/api/v1/Posts/1 --request DELETE

    // get again to check
    curl -u Postname:password localhost:8080/api/v1/Posts
*/


/*
	Non-handlers
*/

type PostFilled struct {
	models.Post // original stuff
	
	Usr models.User `json:"user"`
	Rdg models.Reading `json:"reading"`
}


/*
	Handlers
*/


// this could have a recommendatio algo
// for now most recent.
// populate with 'user' and 'reading'
func GetFeedPosts(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []PostFilled
	// TODO dynamic offset
	// query := "select * from (select * from posts order by updated_at limit 10 offset 0) a join (select id as user_id, email from users) b on a.user_id = b.user_id"

	// get offset from url, set to 0 if none
	offset := req.FormValue("offset")
	if offset == "" { offset = "0" }

	// get posts (most recent first)
	// TODO recommendation logic
	query := "select * from posts order by updated_at desc limit " + POST_OFFSET_INCREMENT + " offset " + offset
	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	// populate users
	var wg sync.WaitGroup
	for i, _ := range items {
		wg.Add(1)

		// item := items[i]
		// change by reference
		go func(i int) {
			defer wg.Done()


			// get user
			user_obj, err := dbmap.Get(models.User{}, items[i].UserId)
			if err != nil { return }
			if user_obj != nil {
				items[i].Usr = *(user_obj.(*models.User)) 
				// fmt.Printf("%v\n", items[i])
			}

			// get reading
			reading_obj, err := dbmap.Get(models.Reading{}, items[i].ReadingId)
			if err != nil { return }
			if reading_obj != nil {
				items[i].Rdg = *(reading_obj.(*models.Reading)) 
				// fmt.Printf("%v\n", items[i])
			}
		}(i)
	}

	wg.Wait()

	// fmt.Printf("%v\n", items)

	Respond(items, err, res)
}





// TODO control viewing based on currentUser and public/private
func GetPosts(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var items []models.Post

	// add query params
	where := addQueryParameters(req, []string{"user_id", "reading_id"})

	query := "select * from posts "+where+" order by updated_at"

	_, err := dbmap.Select(&items, query)
	if err != nil { 
		Respond(nil, err, res)
		return
	}

	Respond(items, err, res)
}

// TODO control viewing based on currentUser and public/private
func GetPost(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.Post
	result, err := getById(item, req, dbmap)
	if err != nil { 
		Respond(nil, err, res)
		return
	}
	Respond(result, err, res)
}



// For another model without hooks (i.e. password->hash), 
// would likely want to use parseFormValues.
func PostPost(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	item := new(models.Post)
	// item.Email = req.FormValue("email")
	// item.SetPassword(req.FormValue("password"))

	err := json.NewDecoder(req.Body).Decode(&item)
	if err != nil { 
		Respond(item, err, res)
		return
	}
	if item.UserId == 0 || item.ReadingId == 0 {
		Respond(item, errors.New("Invalid reading or user id"), res)
		return
	}

	// save Post
	err = dbmap.Insert(item)
	Respond(item, err, res)
}


func PatchPost(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	var item models.Post
	modelname := "posts"

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

func DeletePost(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	DeleteItem("posts", res, req, dbmap)
}
