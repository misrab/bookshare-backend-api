package models


import (
    "time"

    // "github.com/go-gorp/gorp"
    "gopkg.in/gorp.v1"
)



type Post struct {
	

    Comment string `json:"comment"`

    // Associations
    UserId int64 `json:"user_id" db:"user_id"`
    ReadingId int64 `json:"reading_id" db:"reading_id"`


    // default 0 means public innit
    Private bool `json:"private" db:"private"`



    // Meta
    Id int64 `db:"id" json:"id"`
    CreatedAt int64 `db:"created_at" json:"created_at"`
    UpdatedAt int64 `db:"updated_at" json:"updated_at"`
}



/*
 *  SQL Hooks
 */

// implement the PreInsert and PreUpdate hooks
func (i *Post) PreInsert(s gorp.SqlExecutor) error {
    i.CreatedAt = time.Now().UnixNano()
    i.UpdatedAt = i.CreatedAt
    return nil
}

func (i *Post) PreUpdate(s gorp.SqlExecutor) error {
    i.UpdatedAt = time.Now().UnixNano()
    return nil
}

