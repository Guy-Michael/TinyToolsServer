package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Learn GO", Completed: false},
	{ID: "3", Item: "Learn go web apis", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/locations", getData)
	router.POST("/todos", addTodo)
	router.Run("0.0.0.0:8080")
}

func getTodos(context *gin.Context) {
	fmt.Println("GET")
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		context.Error(err)
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getData(context *gin.Context) {
	jsonFile, err := os.Open("data/data.json")

	if err != nil {
		fmt.Println(err)
	}

	context.BindJSON(&jsonFile)
	context.IndentedJSON(http.StatusOK, jsonFile)
}
