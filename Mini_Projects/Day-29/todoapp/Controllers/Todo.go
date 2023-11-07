package Controllers

import (
	"net/http"

	"todoapp/Models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todo []Models.Todo
	err := Models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		// c.JSON(http.StatusOK, todo)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todo,
		})
	}
}

func CreateATodo(c *gin.Context) {
	var todo Models.Todo
	title := c.PostForm("title")
	description := c.PostForm("description")
	todo.Title = title
	todo.Description = description

	if err := Models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, todo)
	c.Redirect(http.StatusSeeOther, "/v1/todo")
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Models.Todo
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)
	err = Models.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
