package handlers
 
import (
	"strings"

    "encoding/base64"
    "net/http"

    _ "github.com/lib/pq"
	"github.com/coopernurse/gorp"
)
 
type handler func(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap)
 
func BasicAuth(pass handler) handler {
    return func(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
 		header := req.Header.Get("Authorization") //req.Header["Authorization"]
 		if header == "" {
 			http.Error(res, "bad syntax", http.StatusBadRequest)
            return
 		}

        auth := strings.SplitN(header, " ", 2)
 
        if len(auth) != 2 || auth[0] != "Basic" {
            http.Error(res, "bad syntax", http.StatusBadRequest)
            return
        }
 
        payload, _ := base64.StdEncoding.DecodeString(auth[1])
        pair := strings.SplitN(string(payload), ":", 2)
 
        if len(pair) != 2 || !Validate(pair[0], pair[1]) {
            http.Error(res, "authorization failed", http.StatusUnauthorized)
            return
        }
 
        pass(res, req, dbmap)
    }
}

func Validate(username, password string) bool {
    if username == "username" && password == "password" {
        return true
    }
    return false
}