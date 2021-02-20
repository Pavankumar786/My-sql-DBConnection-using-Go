package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
type User struct {
    Name string `json:"name"`
}

func main() {
    fmt.Println("Go MySQL")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:Pavan@970@tcp(127.0.0.1:3306)/testdb")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
    results, err :=db.Query("select name from users")
    if err != nil {
        panic(err.Error())
    }
    for results.Next() {
        var user User
        err = results.Scan(&user.Name)
        if err !=nil {
            panic(err.Error())

        }
        fmt.Println(user.Name)
    
}
}