package main

import (
	"context"
	"fmt"
	"github.com/YSecretY/go-rest-api-comments/internal/comment"
	"github.com/YSecretY/go-rest-api-comments/internal/db"
)

func Run() error {
	fmt.Println("starting up...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"9a31bf83-28bc-4b8d-bf70-7d347a24ff2e",
	))
	return nil
}

func main() {
	fmt.Println("Go Rest API Comments")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
