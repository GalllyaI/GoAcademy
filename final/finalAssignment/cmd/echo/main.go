package main

import (
	// "database/sql"
	"final/cmd"
	"final/cmd/echo/handlers"
	"final/cmd/echo/models"

	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "modernc.org/sqlite"
)

func main() {
	router := echo.New()

	db := connectToDB()
	defer db.Close()

	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			log.Println("Echo middleware was called")
			return next(ctx)
		}
	})

	// Add your handler (API endpoint) registrations here
	router.GET("/api/", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})
	router.GET("/api/lists", handlers.HandleGetAllLists)
	router.GET("/api/lists/:id/tasks", handlers.HandleGetTasksByListId)
	router.POST("/api/lists/:id/tasks", handlers.HandleCreateTask)
	router.POST("/api/lists", handlers.HandleCreateList)
	router.PATCH("/api/tasks/:id", handlers.HandleUpdateTask)
	router.DELETE("/api/lists/:id", handlers.HandleDeleteList)
	router.DELETE("/api/tasks/:id", handlers.HandleDeleteTask)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}

func connectToDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite", "data.db")
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to connect: %v", err))
	}
	models.SetDatabase(db)
	db.Exec("PRAGMA foreign_keys = ON;")

	return db

}

// go run cmd/echo/main.go
