package main

import (
    "log"

    "github.com/rongfengliang/ent-demo/ent"
    "context"
    _ "github.com/go-sql-driver/mysql"
)


func main() {
    client, err := ent.Open("mysql", "root:dalongrong@tcp(127.0.0.1)/gogs")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
    defer client.Close()
    // run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}