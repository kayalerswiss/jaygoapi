package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"os"
	"log"

)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{

		"message":"Nattanon Hello All",
	})
}

func main() {

	e := echo.New()
	e.GET("/hello" , helloHandler)

	port := os.Getenv("PORT")
	log.Println("port", port)
	e.Start(":" +port)
	//e.Start(":1323") // list and server on 127.0.0.0:
}