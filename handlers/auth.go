package handlers
 
import (
    // "fmt"
	"strings"

    "encoding/base64"
    "net/http"

    _ "github.com/lib/pq"
	"github.com/go-gorp/gorp"

    "github.com/misrab/bookshare-backend-api/models"
)

const (
    TOKEN = "lalala"
)


 
type handler func(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap)
 

func DecodeAuthHeader(req *http.Request) (string, string, int) {
    header := req.Header.Get("Authorization") //req.Header["Authorization"]
    if header == "" {
        return "", "", http.StatusUnauthorized 
    }

    auth := strings.SplitN(header, " ", 2)

    if len(auth) != 2 || auth[0] != "Basic" {
        println("auth funny")
        return "", "", http.StatusUnauthorized 
    }

    payload, _ := base64.StdEncoding.DecodeString(auth[1])

    pair := strings.SplitN(string(payload), ":", 2)

    

    // TODO check pair len = 2 ?
    if len(pair) < 2 {
        println("pair funny")
        return "", "", http.StatusUnauthorized 
    }

    return pair[0], pair[1], http.StatusOK
    // return pair[0], pair[1], http.StatusOK
}

func BasicAuth(pass handler) handler {
    return func(res http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap) {
        // TEMP
        // pass(res, req, dbmap)
        // return


 		// header := req.Header.Get("Authorization") //req.Header["Authorization"]
 		// if header == "" {
 		// 	http.Error(res, "bad syntax", http.StatusBadRequest)
   //          return
 		// }

   //      auth := strings.SplitN(header, " ", 2)
 
   //      if len(auth) != 2 || auth[0] != "Basic" {
   //          http.Error(res, "bad syntax", http.StatusBadRequest)
   //          return
   //      }
 
   //      payload, _ := base64.StdEncoding.DecodeString(auth[1])
   //      pair := strings.SplitN(string(payload), ":", 2)

        username, hash, errCode := DecodeAuthHeader(req)
        if errCode != http.StatusOK {
            SetHeaders(res, http.StatusUnauthorized)
            http.Error(res, "bad syntax", errCode)
            return
        }
        
        // get the user
        item := new(models.User)
        err := dbmap.SelectOne(&item, "select * from users where email=$1", username)
        if err != nil || item.Hash != hash {
            SetHeaders(res, http.StatusUnauthorized)
            http.Error(res, "bad syntax", errCode)
            return
        }


        // if hash != 
        // if !Validate(username, password) {
        //     println("username password did not validate")
        //     SetHeaders(res, http.StatusUnauthorized)
        //     http.Error(res, "authorization failed", http.StatusUnauthorized)
        //     return
        // }
 
        pass(res, req, dbmap)
    }
}

// func Validate(username, password string) bool {
//     if username == "username" && password == "password" {
//         return true
//     }
//     return false
// }