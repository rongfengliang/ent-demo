package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rongfengliang/ent-demo/ent"
	"github.com/rongfengliang/ent-demo/ent/car"
	"github.com/rongfengliang/ent-demo/ent/group"
	"github.com/rongfengliang/ent-demo/ent/user"
)

func main() {
	client, err := ent.Open("mysql", "root:dalongrong@tcp(127.0.0.1)/gogs")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	CreateGraph(ctx, client)
	QueryGithub(ctx, client)
	QueryArielCars(ctx, client)
	QueryGroupWithUsers(ctx, client)
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// first, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// then, create the cars, and attach them to the users in the creation.
	_, err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta).              // attach this graph to Neta.
		Save(ctx)
	if err != nil {
		return err
	}
	// create the groups, and add their users in the creation.
	_, err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.
		Query().
		Where(group.Name("GitHub")). // (Group(Name=GitHub),)
		QueryUsers().                // (User(Name=Ariel, Age=30),)
		QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %v", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
	return nil
}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
	// Get "Ariel" from previous steps.
	a8m := client.User.
		Query().
		Where(
			user.HasCars(),
			user.Name("Ariel"),
		).
		OnlyX(ctx)
	cars, err := a8m. // Get the groups, that a8m is connected to:
				QueryGroups(). // (Group(Name=GitHub), Group(Name=GitLab),)
				QueryUsers().  // (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
				QueryCars().   //
				Where(         //
			car.Not( //  Get Neta and Ariel cars, but filter out
				car.ModelEQ("Mazda"), //  those who named "Mazda"
			), //
		). //
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting cars: %v", err)
	}
	log.Println("cars returned:", cars)
	// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
	return nil
}
func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.
		Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed getting groups: %v", err)
	}
	log.Println("groups returned:", groups)
	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
	return nil
}
