package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Junkes887/sqlc/internal/database"
)

func main() {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer con.Close()
	db := database.New(con)
	ctx := context.Background()

	db.CreateUser(ctx, database.CreateUserParams{ID: 1, Name: "John"})
	db.CreateIten(ctx, database.CreateItenParams{Name: "pen", UserID: 1})

	itens, _ := db.GetItens(ctx)
	for _, iten := range itens {
		user, _ := db.GetUser(ctx, iten.UserID)
		fmt.Printf("Item id: %d, Item name: %s, User id: %d, User name: %s \n", iten.ID, iten.Name, user.ID, user.Name)
	}
}
