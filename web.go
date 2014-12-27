package main 

import(
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/misrab/goapi/handlers"
    "github.com/misrab/goapi/models"
)




func main() {
    router := mux.NewRouter()
    dbmap := models.SetupDB()

    
    /*
        Routes
    */

    // User
    // router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
    //     handlers.BasicAuth(handlers.PostUser(res, req, dbmap)
    // }).Methods("POST")
    // router.HandleFunc("/api/v1/users/{id}", func(res http.ResponseWriter, req *http.Request) {
    //     handlers.PatchUser(res, req, dbmap)
    // }).Methods("PATCH")
    router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
        //handlers.GetUsers(res, req, dbmap)
        handlers.BasicAuth(handlers.GetUsers)(res, req, dbmap)
    }).Methods("GET")
    // // general
    // router.HandleFunc("/api/v1/{resource}/{id}", func(res http.ResponseWriter, req *http.Request) {
    //     var i models.User
    //     handlers.GetItem(i, res, req, dbmap)
    // }).Methods("GET")
    // // general
    // router.HandleFunc("/api/v1/users/{id}", func(res http.ResponseWriter, req *http.Request) {
    //     handlers.DeleteItem("users", res, req, dbmap)
    // }).Methods("DELETE")
    

    
    // register routes
    http.Handle("/", router)

    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}