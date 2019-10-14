package main

import (
	"context"
	"fmt"
	"github.com/rongfengliang/ent-demo/ent"
	"github.com/rongfengliang/ent-demo/ent/user"
	"log"
)

func main()  {

	client, err := ent.Open("mysql", "root:dalongrong@tcp(127.0.0.1)/gogs")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	QueryUser(ctx, client)
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.ID(1)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
