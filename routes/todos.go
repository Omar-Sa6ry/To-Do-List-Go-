package routes

import (
	"net/http"
	"strconv"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func getTodos(context *gin.Context) {
	userId := context.GetInt64("userId")
	
	todos, err := models.GetAllTodos(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todos. Try again later."})
		return
	}

	context.JSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id"})
		return
	}

	userId := context.GetInt64("userId")

	todo, err := models.GetTodoByID(id, userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.JSON(http.StatusOK, todo)
}

func createTodo(context *gin.Context) {
	var todo models.Todo
	err := context.ShouldBindJSON(&todo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	todo.UserID = context.GetInt64("userId")

	err = todo.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create todo. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Todo created!", "todo": todo})
}

func updateTodo(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id"})
		return
	}

	userId := context.GetInt64("userId")

	_, err = models.GetTodoByID(id, userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	var updatedTodo models.Todo
	err = context.ShouldBindJSON(&updatedTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedTodo.ID = id
	updatedTodo.UserID = userId

	err = updatedTodo.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update todo"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully!"})
}

func deleteTodo(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo id"})
		return
	}

	userId := context.GetInt64("userId")

	_, err = models.GetTodoByID(id, userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	err = models.DeleteTodo(id, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete todo"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully!"})
}
