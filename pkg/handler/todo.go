package handler

import (
	todo "crud"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllTodosResponse struct {
	Result []todo.Todo `json:"result"`
}

func (h *Handler) getAllTodos(c *gin.Context) {
	todos, err := h.services.Todo.GetAllTodo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllTodosResponse{
		Result: todos,
	})
}

func (h *Handler) createTodo(c *gin.Context) {
	var input todo.Todo

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Todo.CreateTodo(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) changeIsDone(c *gin.Context) {
	var input struct {
		IsDone bool `json:"isDone,omitempty"`
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	_, err = h.services.Todo.ChangeDoneStatus(todo.ChangeDoneStatusDto{Id: id, IsDone: input.IsDone})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) changeIsFavourite(c *gin.Context) {
	var input struct {
		IsFavourite bool `json:"isFavourite,omitempty"`
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	_, err = h.services.Todo.ChangeFavouriteStatus(todo.ChangeFavouriteStatusDto{Id: id, IsFavourite: input.IsFavourite})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	_, err = h.services.Todo.DeleteTodo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
