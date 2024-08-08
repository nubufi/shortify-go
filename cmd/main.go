package main

import (
	"shortify/lib"
	"shortify/web/backend"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := lib.CreateDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := echo.New()

	app.Static("/static", "static")
	app.GET("/", backend.HomePageHandler)
	app.GET("/:shortCode", func(c echo.Context) error {
		return backend.RedirectHandler(c, db)
	})
	app.POST("/shorten", func(c echo.Context) error {
		return backend.ShortenHandler(c, db)
	})

	app.Logger.Fatal(app.Start(":8080"))
}
