package scraper

import (
    // "bufio"
    "fmt"
    //    "strings"
    //    "io"
    // "log"
    // "os"
    // "time"



    "encoding/json"


    // "database/sql"
    // _ "github.com/lib/pq"
    // "github.com/coopernurse/gorp"


    "github.com/misrab/bookshare-backend-api/models"
    
    "github.com/misrab/goutils"
)


const (

)

// TODO get open library dumps
// see https://openlibrary.org/developers/dumps
// test in /data/dumps_short_works.txt...


type BookWrapper struct {
    models.Reading

    Covers []int `json:"covers"`
}

func GetBooks() {
	println("getting books")


    filepath := "./data/dumps_short_works.txt"


    // connection to database
    dbmap := models.ConnectDB()
    

    // read from dump
    c := make(chan []string)

    go utils.ReadCsv(filepath, c, true, '\t')

    for record := range c {
        

        // book := new(models.Reading)
        book := new(BookWrapper)

        // TODO make a suitable object
        json.Unmarshal([]byte(record[4]), book)

        // other fields
        book.IsBook = true



        // go from wrapper to actual db model
        if len(book.Covers) > 0 { 
            book.Reading.Cover = book.Covers[0]
        }
        // book.Reading.Cover = 222
        // reading := new(models.Reading)
        reading := &book.Reading



        fmt.Printf("%+v\n", reading)
        // add to database
        err := dbmap.Insert(reading)
        if err != nil { panic(err) }

        // TEMP
        fmt.Printf("%s\n", "finished inserting")
        return

    }

    // TODO resolve authors or store them


    
}




// convenience debug method
func readDump(filepath string) {
    c := make(chan []string)

    go utils.ReadCsv(filepath, c, true, '\t')

    for record := range c {
        fmt.Printf("%v\n", record)
    }
}