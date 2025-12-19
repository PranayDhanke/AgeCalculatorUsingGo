package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/pranaydhanke/go-user-api/config"
	"github.com/pranaydhanke/go-user-api/db/sqlc"
	"github.com/pranaydhanke/go-user-api/internal/handler"
	"github.com/pranaydhanke/go-user-api/internal/logger"
	"github.com/pranaydhanke/go-user-api/internal/repository"
	"github.com/pranaydhanke/go-user-api/internal/routes"
)

func main() {
	cfg := config.Load()
	logg := logger.New()
	defer logg.Sync()

	db, err := sql.Open("pgx", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	queries := sqlc.New(db)
	repo := repository.NewUserRepository(queries)
	h := handler.NewUserHandler(repo)

	app := fiber.New()
	routes.Register(app, h)

	app.Use(func(c *fiber.Ctx) error {
		println("Incoming:", c.Method(), c.Path())
		return c.Next()
	})

	log.Fatal(app.Listen(":8080"))
}
