package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"log"

)

type Todo struct {
	id int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status:"active"},
}

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{

		"message":"Nattanon Hello All",
	})
}

func getTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}

}

func main() {

	e := echo.New()

	//Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello" , helloHandler)
	e.GET("/todos",getTodosHandler)

	port := os.Getenv("PORT")
	log.Println("port", port)
	e.Start(":" +port)
	//e.Start(":1323") // list and server on 127.0.0.0:
}