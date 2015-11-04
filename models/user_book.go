package models


import (
    "time"

    "github.com/go-gorp/gorp"
)


// join table
type UserBook struct {
	

    UserId int64 `db:"user_id"`
    BookId int64 `db:"book_id"`


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
func (i *UserBook) PreInsert(s gorp.SqlExecutor) error {
    i.CreatedAt = time.Now().UnixNano()
    i.UpdatedAt = i.CreatedAt
    return nil
}

func (i *UserBook) PreUpdate(s gorp.SqlExecutor) error {
    i.UpdatedAt = time.Now().UnixNano()
    return nil
}
