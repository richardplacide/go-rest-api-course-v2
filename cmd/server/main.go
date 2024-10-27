package main

import (
	"context"
	"fmt"

	"orbitalcoding.com/go-rest-api-course-v2/internal/db"
)

func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("could not connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		fmt.Println("could not ping the database")
		return err
	}

	return nil

}

func main() {
	fmt.Println("Hello, World!")
	if err := Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
