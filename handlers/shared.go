package handlers


import (
	// "log"
	// "time"
	"strconv"
	// "errors"
	
	"net/http"
	// "encoding/json"

	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"
	
	"github.com/gorilla/mux"
	// "github.com/gorilla/schema"

	//"github.com/misrab/minum/models"
)

/*
	All
*/



/*
	Associations
*/




/*
	Get
*/

func getId(req *http.Request) (int64, error) {
	vars := mux.Vars(req)
	return strconv.ParseInt(vars["id"], 0, 64)
}

func getById(i interface{}, req *http.Request, dbmap *gorp.DbMap) (interface{}, error) {
	// vars := mux.Vars(req)
	// id, err1 := strconv.ParseInt(vars["id"], 0, 64)
	id, err1 := getId(req)
	if err1 != nil { return nil, err1 }
	obj, err2 := dbmap.Get(i, id)
	
	return obj, err2
}


/*
	Patch
*/

func parseFormValues(i interface{}, req *http.Request) error {
	var err error
	err = req.ParseForm()
	if err != nil { return err }

	err = decoder.Decode(i, req.PostForm)	
	//err = json.NewDecoder(req.Body).Decode(i)

	return err
}


/*
	Delete
*/

func DeleteItem(modelname string, res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
	vars := mux.Vars(req)
	id := vars["id"]
	query := "delete from " + modelname + " where id=" + id
	_, err := dbmap.Exec(query)

	Respond(nil, err, res)
}
