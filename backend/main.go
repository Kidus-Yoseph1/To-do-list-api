package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var lists = []Todo{
	{ID: 1, Title: "Assignment", Completed: false, CreatedAt: time.Now()},
	{ID: 2, Title: "Practice GO", Completed: true, CreatedAt: time.Now()},
}

func getTodo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, lists)
}

func creatTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if strings.TrimSpace(newTodo.Title) == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title can not be empty"})
	}
	newTodo.ID = len(lists) + 1
	lists = append(lists, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	id_param := c.Param("id")
	id, _ := strconv.Atoi(id_param)
	for i, todo := range lists {
		if todo.ID == id {
			lists[i].Completed = true
			c.JSON(http.StatusOK, lists[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
}

func delete(c *gin.Context) {
	id_param := c.Param("id")
	id, _ := strconv.Atoi(id_param)
	for i, todo := range lists {
		if todo.ID == id {
			lists = append(lists[:i], lists[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
}

func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/lists", getTodo)
		api.POST("/lists", creatTodo)
		api.PATCH("/update/:id", updateTodo)
		api.DELETE("/tasks/:id", delete)
	}
	r.Run("localhost:8080")
}
