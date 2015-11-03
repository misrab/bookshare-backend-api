package main 

import(
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/misrab/bookshare-backend-api/handlers"
    "github.com/misrab/bookshare-backend-api/models"
)



func main() {
    router := mux.NewRouter()
    dbmap := models.SetupDB()

    
    /*
        Routes
    */

    // User
    router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetUsers)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/users/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetUser)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.PostUser)(res, req, dbmap)
    }).Methods("POST")
    router.HandleFunc("/api/v1/users/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.PatchUser)(res, req, dbmap)
    }).Methods("PATCH")
    router.HandleFunc("/api/v1/users/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.DeleteUser)(res, req, dbmap)
    }).Methods("DELETE")

    // router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
    //     handlers.BasicAuth(handlers.PostUser(res, req, dbmap)
    // }).Methods("POST")
    
    // // general
    
    // // general
    
    

    
    // register routes
    http.Handle("/", router)

    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}