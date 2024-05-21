package main

import (
	"context"
	"database/sql"
	"github.com/HallZero/go-frameworks/sqlc/internal/db"
)

func main() {

	dbconn, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")

	if err != nil {
		panic(err.Error())
	}

	defer dbconn.Close()

	database := db.New(dbconn)

	ctx := context.Background()
	
	err = database.CreateCategory(ctx, db.CreateCategoryParams{
		ID: 1,
		Name: "Test",
	})

	if err != nil {
		panic(err.Error())
	}
}