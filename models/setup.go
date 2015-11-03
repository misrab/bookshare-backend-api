package models

import (
    "os"
    "fmt"
    "log"
    "regexp"

    "database/sql"

    _ "github.com/lib/pq"
    "github.com/coopernurse/gorp"
)

//func SetupDB() *sql.DB {
func SetupDB() *gorp.DbMap {
    // get environment
    env := os.Getenv("ENV")
    println('env is ' + env)

    db := pgConnect()

    // construct a gorp DbMap
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

    // add a table, setting the table name to 'posts' and
    // specifying that the Id property is an auto incrementing PK
    dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id").ColMap("Email").SetUnique(true)
    // dbmap.AddTableWithName(Topic{}, "topics").SetKeys(true, "Id")
    // dbmap.AddTableWithName(UserTopic{}, "user_topics")
    // dbmap.AddTableWithName(UserTopic{}, "user_quests")

    // dbmap.AddTableWithName(Resource{}, "resources").SetKeys(true, "Id")
    // dbmap.AddTableWithName(Quest{}, "quests").SetKeys(true, "Id")
    // dbmap.AddTableWithName(Discussion{}, "discussions").SetKeys(true, "Id")
    // dbmap.AddTableWithName(Comment{}, "comments").SetKeys(true, "Id")


    // drop all tables for testing
    log.Println("DROPPING TABLES!")
    err1 := dbmap.DropTablesIfExists()
    if err1 != nil { panic(err1) }

    err2 := dbmap.CreateTablesIfNotExists()
    // PanicIf(err2)
    if err2 != nil {
        panic(err2)
    }

    // set logging for development
    dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds)) 
    
    return dbmap
}


func pgConnect() *sql.DB {
    // Connect to Postgres database
    env := os.Getenv("ENV")
    regex := regexp.MustCompile("(?i)^postgres://(?:([^:@]+):([^@]*)@)?([^@/:]+):(\\d+)/(.*)$")
    var connection string
    switch env {
    //case "staging":
    //case "production":
    // default to development
    default:
        connection = os.Getenv("POSTGRESQL_LOCAL_URL")
    }
    matches := regex.FindStringSubmatch(connection)
    sslmode := "disable"
    spec := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", matches[1], matches[2], matches[3], matches[4], matches[5], sslmode)

    db, err := sql.Open("postgres", spec)
    //PanicIf(err)
    if err != nil {
        panic(err)
    }

    log.Printf("Connected to %s\n", connection)

    return db
}