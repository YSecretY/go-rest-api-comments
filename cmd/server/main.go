package main

import (
	"fmt"
	"github.com/YSecretY/go-rest-api-comments/internal/comment"
	"github.com/YSecretY/go-rest-api-comments/internal/db"
	transportHttp "github.com/YSecretY/go-rest-api-comments/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest API Comments")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
