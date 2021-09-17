package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/shnaki/ent-example/ent"
	"github.com/shnaki/ent-example/ent/user"
)

func main() {
	client, err := ent.Open("sqlite3",
		"file:./ent.db?_loc=auto&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	u, err := CreateUser(context.Background(), client)
	if err != nil {
		log.Fatalf("failed creating user table: %v", err)
	}
	fmt.Println(u)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
