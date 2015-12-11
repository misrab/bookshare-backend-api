package models


import (
    "time"

    // "golang.org/x/crypto/bcrypt"
    "code.google.com/p/go.crypto/bcrypt"
    "github.com/go-gorp/gorp"
)



type User struct {
	Id       	int64 `db:"id" json:"id"`
    

    Email  		string `json:"email"` //`schema:"email",json:"email",db:"email"`
    Hash 		string `json:"hash"`

    CreatedAt int64 `db:"created_at" json:"created_at"`
    UpdatedAt int64 `db:"updated_at" json:"updated_at"`
}




func (user *User) SetPassword(password string) {
    // use negative cost to prompt default
    hash, err := bcrypt.GenerateFromPassword([]byte(password) , -1)

    if err != nil { return }
    user.Hash = string(hash)
}

// Returns nil error on sucess
func (user *User) ComparePassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
}


/*
 *  SQL Hooks
 */

// implement the PreInsert and PreUpdate hooks
func (i *User) PreInsert(s gorp.SqlExecutor) error {
    i.CreatedAt = time.Now().UnixNano()
    i.UpdatedAt = i.CreatedAt
    return nil
}

func (i *User) PreUpdate(s gorp.SqlExecutor) error {
    i.UpdatedAt = time.Now().UnixNano()
    return nil
}