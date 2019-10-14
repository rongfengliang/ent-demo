package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rongfengliang/ent-demo/ent"
	"log"
)
func main() {
	client, err := ent.Open("mysql", "root:dalongrong@tcp(127.0.0.1)/gogs")
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }
	defer client.Close()
	ctx := context.Background()
	CreateUser(ctx,client)

}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Create().
        SetAge(30).
        SetName("a8m").
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %v", err)
    }
    log.Println("user was created: ", u)
    return u, nil
}