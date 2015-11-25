package models


import (
    "time"

    "github.com/go-gorp/gorp"
)



type Reading struct {
	

    Title string `json:"title"`




    // Fields found so far
    // subtitle
    // covers
    // latest_revision, revision

    // subject_times
    // subjects

    IsBook bool `json:"is_book"`



    // TODO Authors - likely a join table

    // Meta
    Id int64 `db:"id" json:"id"`
    CreatedAt int64 `db:"created_at" json:"created_at"`
    UpdatedAt int64 `db:"updated_at" json:"updated_at"`
}



/*
 *  SQL Hooks
 */

// implement the PreInsert and PreUpdate hooks
func (i *Reading) PreInsert(s gorp.SqlExecutor) error {
    i.CreatedAt = time.Now().UnixNano()
    i.UpdatedAt = i.CreatedAt
    return nil
}

func (i *Reading) PreUpdate(s gorp.SqlExecutor) error {
    i.UpdatedAt = time.Now().UnixNano()
    return nil
}

