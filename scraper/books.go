package scraper

import (
    // "bufio"
    "fmt"
    //    "strings"
    //    "io"
    // "log"
    "strconv"
    "os"
    "time"


    "encoding/csv"
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
    // dbmap := models.ConnectDB()


    // file to write to
    f, err := os.Create("out.csv")
    if err != nil { panic(err) }
    defer f.Close()
    writer := csv.NewWriter(f)

    
    now := strconv.FormatInt(time.Now().UnixNano(), 10)

    // read from dump
    c := make(chan []string)

    go utils.ReadCsv(filepath, c, true, '\t')

    count := 1
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



        new_row := []string{
            reading.Title, // title
            reading.Key,   // key
            strconv.Itoa(reading.Cover), // cover []
            "true", // isbook
            strconv.Itoa(count), // id
            now,
            now,
        }

        // fmt.Printf("%+v\n", reading)
        // fmt.Printf("writing %+v\n", record)

        err := writer.Write(new_row)
        if err != nil { panic(err) }


        if (count%100 == 0) { 
            println("Flushing row " + strconv.Itoa(count))
            writer.Flush()
        }

        // add to database
        // err := dbmap.Insert(reading)
        // if err != nil { 
        //     // continue
        //     panic(err) 
        // }

        // return

        count++

        // break
    }


    writer.Flush()
    

    // TODO resolve authors or store them


    
}


// COPY readings FROM '/Users/Misrab/go/src/github.com/misrab/bookshare-backend-api/scraper/out.csv' DELIMITER ',' CSV;
// psql -h <host> -p <port> -u <database>
// psql -h <host> -p <port> -U <username> -W <password> <database>
// psql -h booksharepsql.cqhcjpglhfga.ap-southeast-1.rds.amazonaws.com -p 5432 -U misrab -W postgres

// convenience debug method
func readDump(filepath string) {
    c := make(chan []string)

    go utils.ReadCsv(filepath, c, true, '\t')

    for record := range c {
        fmt.Printf("%v\n", record)
    }
}