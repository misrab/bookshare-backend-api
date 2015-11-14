package main 

import(
	"log"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/misrab/bookshare-backend-api/handlers"
    "github.com/misrab/bookshare-backend-api/models"
)


// Simple wrapper to Allow CORS.
// func withCORS(fn handlers.handler) http.HandlerFunc {
//  return func(w http.ResponseWriter, r *http.Request) {
//    w.Header().Set("Access-Control-Allow-Origin", "*")
//    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//    fn(w, r)
//  }
// }

// func corsHandler(h http.Handler) http.HandlerFunc {
//   return func(w http.ResponseWriter, r *http.Request) {
//     if (r.Method == "OPTIONS") {
//       //handle preflight in here
//     } else {
//       h.ServeHTTP(w,r)
//     }
//   }
// }

func main() {
    router := mux.NewRouter()
    dbmap := models.SetupDB()

    
    /*
        Routes
    */


    // CORS preflight request
    router.HandleFunc("/api/v1/{*}", func(res http.ResponseWriter, req *http.Request) {
    // handlers.Respond(nil, nil, res)
        // handlers.SetHeaders(res, 200)
        // res.WriteHeader(200)
        handlers.Respond(nil, nil, res)
    // res.Write(nil)
    }).Methods("OPTIONS")
    // router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
    // // handlers.Respond(nil, nil, res)
    //     // handlers.SetHeaders(res, 200)
    //     // res.WriteHeader(200)
    //     handlers.Respond(nil, nil, res)
    // // res.Write(nil)
    // }).Methods("OPTIONS")


    // User
    router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetUsers)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/users/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetUser)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/users", func(res http.ResponseWriter, req *http.Request) {
        // handlers.BasicAuth(handlers.PostUser)(res, req, dbmap)
        handlers.PostUser(res, req, dbmap)
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
    // get current user based on Authorization header
    router.HandleFunc("/api/v1/current_user", func(res http.ResponseWriter, req *http.Request) {
        // handlers.BasicAuth(handlers.GetCurrentUser)(res, req, dbmap)
        handlers.GetCurrentUser(res, req, dbmap)
    }).Methods("GET")
    // router.HandleFunc("/api/v1/current_user", func(res http.ResponseWriter, req *http.Request) {
    //     res.Header().Set("Access-Control-Allow-Origin", "*")
    //     res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    //     handlers.BasicAuth(handlers.GetCurrentUser)(res, req, dbmap)
    // }).Methods("OPTIONS")
    router.HandleFunc("/api/v1/login", func(res http.ResponseWriter, req *http.Request) {
        // handlers.BasicAuth(handlers.LoginUser)(res, req, dbmap)
        handlers.LoginUser(res, req, dbmap)
    }).Methods("POST")


    // Reading
    router.HandleFunc("/api/v1/readings", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetReadings)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/readings/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetReading)(res, req, dbmap)
    }).Methods("GET")
    router.HandleFunc("/api/v1/readings", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.PostReading)(res, req, dbmap)
    }).Methods("POST")
    router.HandleFunc("/api/v1/readings/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.PatchReading)(res, req, dbmap)
    }).Methods("PATCH")
    router.HandleFunc("/api/v1/readings/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.DeleteReading)(res, req, dbmap)
    }).Methods("DELETE")
    // Reading autocomplete
    router.HandleFunc("/api/v1/readings_autocomplete", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.GetReadingsAutocomplete)(res, req, dbmap)
    }).Methods("GET")
    

    // User-Reading Association
    // router.HandleFunc("/api/v1/users_Readings", func(res http.ResponseWriter, req *http.Request) {
    //     handlers.BasicAuth(handlers.GetReadings)(res, req, dbmap)
    // }).Methods("GET")
    // router.HandleFunc("/api/v1/users_Readings/{id}", func(res http.ResponseWriter, req *http.Request) {
    //     handlers.BasicAuth(handlers.GetReading)(res, req, dbmap)
    // }).Methods("GET")
    router.HandleFunc("/api/v1/users_readings", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.PostUserReading)(res, req, dbmap)
    }).Methods("POST")
    // router.HandleFunc("/api/v1/Readings/{id}", func(res http.ResponseWriter, req *http.Request) {
    //     handlers.BasicAuth(handlers.PatchReading)(res, req, dbmap)
    // }).Methods("PATCH")
    router.HandleFunc("/api/v1/users_readings/{id}", func(res http.ResponseWriter, req *http.Request) {
        handlers.BasicAuth(handlers.DeleteUserReading)(res, req, dbmap)
    }).Methods("DELETE")


    
    

    
    // register routes
    http.Handle("/", router)

    log.Println("Listening...")
    http.ListenAndServe(":8000", nil)
}