package models

import (
    "os"
    "fmt"
    "log"
    "regexp"

    "database/sql"
    _ "github.com/lib/pq"
    // "github.com/go-gorp/gorp"

    "gopkg.in/gorp.v1"
)

func ConnectDB() *gorp.DbMap {
    db := pgConnect()
    env := os.Getenv("ENV")

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    setupDbmapTables(dbmap)

    if env == "development" {
        // set logging for development
        dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds)) 
    }

    return dbmap
}

// psql -h <host> -p <port> -u <database>


func setupDbmapTables(dbmap *gorp.DbMap) {
    dbmap.AddTableWithNameAndSchema(User{}, "public", "users").SetKeys(true, "Id").ColMap("Email").SetUnique(true)
    // // TODO check if this title index is right
    dbmap.AddTableWithNameAndSchema(Reading{}, "public", "readings").SetKeys(true, "Id") // .AddIndex("ReadingsIndex", "Btree", []string{"Title"}) // .ColMap("Title").SetUnique(true)
    dbmap.AddTableWithNameAndSchema(Post{}, "public", "posts").SetKeys(true, "Id")
    dbmap.AddTableWithNameAndSchema(UserReading{}, "public", "users_readings").SetUniqueTogether("user_id", "reading_id") // join table
}

//func SetupDB() *sql.DB {
func SetupDB() *gorp.DbMap {
    // get environment
    env := os.Getenv("ENV")
    if (env != "development" &&
        env != "staging" &&
        env != "production" ) {
        env = "production" // pick most conservative by default
    }

    
    db := pgConnect()

    // construct a gorp DbMap
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}


    setupDbmapTables(dbmap)

    // log.Println("DROPPING TABLES!")
    // err := dbmap.DropTablesIfExists()
    // if err != nil { panic(err) }

    // drop all tables for testing
    if env == "development" {
        // log.Println("DROPPING TABLES!")
        // err := dbmap.DropTablesIfExists()
        // if err != nil { panic(err) }

        // set logging for development
        dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds)) 
    }
  
    fmt.Printf("Env is %s...\n", env)
    log.Println("Creating tables...")
    err := dbmap.CreateTablesIfNotExists()
    // fmt.Printf("err is %v...\n", err)
    if err != nil {
        panic(err)
    }

    return dbmap
}


func pgConnect() *sql.DB {
    // Connect to Postgres database
    env := os.Getenv("ENV")
    regex := regexp.MustCompile("(?i)^postgres://(?:([^:@]+):([^@]*)@)?([^@/:]+):(\\d+)/(.*)$")
    
    sslmode := "disable"
    var connection string
    switch env {
    //case "staging":
    //case "production":
    // default to development
    default:
        connection = os.Getenv("POSTGRESQL_LOCAL_URL")
    }

    // connection = os.Getenv("POSTGRESQL_LOCAL_URL")
    
    matches := regex.FindStringSubmatch(connection)
    

    // fmt.Printf("%v\n", matches)

    spec := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", matches[1], matches[2], matches[3], matches[4], matches[5], sslmode)

    db, err := sql.Open("postgres", spec)
    //PanicIf(err)
    if err != nil {
        panic(err)
    }

    log.Printf("Connected to %s\n", connection)

    return db
}