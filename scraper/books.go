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


func GetBooks() {
	println("getting books")


    filepath := "./data/dumps_short_works.txt"
	// readDump(filepat)


    // connection to database
    // dbmap := models.ConnectDB()
    dbmap := models.SetupDB()
    
    // TODO get dump and decompress

    // read from dump
    c := make(chan []string)

    go utils.ReadCsv(filepath, c, true, '\t')

    for record := range c {
        

        book := new(models.Book)

        // TODO make a suitable object
        json.Unmarshal([]byte(record[4]), book)
        fmt.Printf("%v\n", book)

        // TODO add to database
        // err := 
        dbmap.Insert(book)


        // TEMP
        // return

    }

    // TODO resolve authors or store them


    
}





func readDump(filepath string) {
    c := make(chan []string)

    go utils.ReadCsv(filepath, c, true, '\t')

    for record := range c {
        fmt.Printf("%v\n", record)
    }

	// file, err := os.Open(filepath)
 //    if err != nil {
 //        log.Fatal(err)
 //    }
 //    defer file.Close()

 //    scanner := bufio.NewScanner(file)
 //    for scanner.Scan() {
 //        fmt.Println(scanner.Text())
 //        return
 //    }

 //    if err := scanner.Err(); err != nil {
 //        log.Fatal(err)
 //    }
}