package main

import (
	"Mini-project/database"
	"Mini-project/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	database.DBMigration()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8090"))
}
