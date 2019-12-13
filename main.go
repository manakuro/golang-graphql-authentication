package main

import (
	"golang-graphql-authentication/auth"
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
	e.POST("/signIn", handler.SignIn())

	// Restricted from here
	r := e.Group("graphql")

	key, err := auth.GetRSAPublicKey()
	logFatal(err)

	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    key,
		SigningMethod: "RS256",
	}))

	// graphql
	h, err := graphql.NewHandler(db)
	logFatal(err)
	r.POST("", echo.WrapHandler(h))

	err = e.Start(":3000")
	logFatal(err)
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
