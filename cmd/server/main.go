package main

import (
	"context"
	"fmt"
	"github.com/YSecretY/go-rest-api-comments/internal/db"
)

func Run() error {
	fmt.Println("starting up...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully connected and pined database")
	return nil
}

func main() {
	fmt.Println("Go Rest API Comments")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
