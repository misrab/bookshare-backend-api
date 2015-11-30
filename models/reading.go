package models


import (
    "time"

    "github.com/go-gorp/gorp"
)



type Reading struct {
	
    // Open library fields
    // Fields found so far
    // subtitle
    // covers
    // latest_revision, revision

    // subject_times
    // subjects
    Title string `json:"title"`

    // OL key e.g. "/works/OL10000223W"
    Key string `json:"key"`
    // cover ids e.g. [3140972]
    // get at http://covers.openlibrary.org/b/id/{id}-[S|M|L].jpg
    Cover int `json:"cover"` //db:"covers,size:1024"



    // TODO Authors - likely a join table


    // our extra fields
    IsBook bool `json:"is_book"`

    
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

