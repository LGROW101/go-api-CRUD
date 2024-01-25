package main

import (
	"go-clouds/pkg/db"
	"go-clouds/pkg/handlers"
	"go-clouds/pkg/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()

	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.JsonMiddleware)

	// Routes
	handlers.SetupRoutes(router, database)

	log.Fatal(http.ListenAndServe(":8000", router))
}
