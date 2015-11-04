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


    // Book
    router.HandleFunc("/api/v1/books", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetBooks)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/books/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetBook)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/books", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.PostBook)(res, req, dbmap)
    }).Methods("POST")
    router.HandleFunc("/api/v1/books/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.PatchBook)(res, req, dbmap)
    }).Methods("PATCH")
    router.HandleFunc("/api/v1/books/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.DeleteBook)(res, req, dbmap)
    }).Methods("DELETE")
    // Book autocomplete
    router.HandleFunc("/api/v1/books_autocomplete", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetBooksAutocomplete)(res, req, dbmap)
    }).Methods("GET")
    

    
    

    
    // register routes
    http.Handle("/", router)

    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}