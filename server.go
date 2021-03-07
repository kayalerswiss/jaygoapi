package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"log"

)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status:"active"},
	2: &Todo{ID: 2, Title: "pay credit card", Status:"inactive"},
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
	return c.JSON(http.StatusOK,items)
}

func createTodosHandler(e echo.Context) error{
	t := Todo{}
	if err := e.Bind(&t); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t

	return e.JSON(http.StatusCreated, "create todo")
}

func getTodoByIdHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id",&id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	t,ok := todos[id]
	if !ok {
		return c.JSON(http.StatusOK, map[int]string{})
	}
	return c.JSON(http.StatusOK,t)
}

// ส่งค่าแบบ PUT (Update)
func updateTodoByIdHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id",&id).BindError()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	t,ok := todos[id]
	if !ok {
		return c.JSON(http.StatusOK, map[int]string{})
	}
	return c.JSON(http.StatusOK,t)
}

func main() {

	e := echo.New()

	//Middleware
	e.Use(middleware.Logger()) // แสดง Log ใน heroku ทุกครั้งที่มีการยิง
	e.Use(middleware.Recover()) // ไม่ให้ server ดับ

	e.GET("/hello" , helloHandler)
	e.GET("/todos",getTodosHandler)
	e.GET("/todos/:id",getTodoByIdHandler)
	e.POST("/todos",createTodosHandler)



	e.PUT("/todos/:id",updateTodoByIdHandler)


	port := os.Getenv("PORT")
	log.Println("port", port)
	e.Start(":" +port)
	//e.Start(":1323") // list and server on 127.0.0.0:
}