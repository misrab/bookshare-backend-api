package models


import (
    "time"

    "github.com/go-gorp/gorp"
)



type Book struct {
	

    Title string `json:"title"`



    // Fields found so far
    // subtitle
    // covers
    // latest_revision, revision

    // subject_times
    // subjects



    // TODO Authors - likely a join table

    // Meta
    Id int64 `db:"id"`
    CreatedAt int64 `db:"created_at"`
    UpdatedAt int64 `db:"updated_at"`
}



/*
 *  SQL Hooks
 */

// implement the PreInsert and PreUpdate hooks
func (i *Book) PreInsert(s gorp.SqlExecutor) error {
    i.CreatedAt = time.Now().UnixNano()
    i.UpdatedAt = i.CreatedAt
    return nil
}

func (i *Book) PreUpdate(s gorp.SqlExecutor) error {
    i.UpdatedAt = time.Now().UnixNano()
    return nil
}

