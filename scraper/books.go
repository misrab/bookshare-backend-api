package scraper

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


const (

)

// TODO get open library dumps
// see https://openlibrary.org/developers/dumps
// test in /data/dumps_short_works.txt...


func GetBooks() {
	println("getting books")

	readDump("./data/dumps_short_works.txt")
}


func readDump(filepath string) {
	file, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        return
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}