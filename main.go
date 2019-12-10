package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"golang-graphql-authentication/datastore"
	"golang-graphql-authentication/graphql"
	"golang-graphql-authentication/handler"
)

func main() {
	db, err := datastore.NewDB()
	logFatal(err)

	db.LogMode(true)
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", handler.Welcome())

	// graphql
	h, err := graphql.NewHandler(db)
	logFatal(err)
	e.POST("/graphql", echo.WrapHandler(h))

	err = e.Start(":3000")
	logFatal(err)
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
