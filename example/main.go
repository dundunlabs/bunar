package main

import (
	"database/sql"
	"os"

	"github.com/dundunlabs/bunar"
	_ "github.com/joho/godotenv/autoload"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type User struct {
}

type Model struct {
	db *bun.DB
}

func (m Model) User() *bunar.Model[User] {
	return bunar.NewModel(m.db, User{})
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	model := Model{db: db}

	model.User()
}
