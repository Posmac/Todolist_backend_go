package handler

import (
	todo "Notes_GoRest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUsersId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.TodoList

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUsersId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUsersId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	//UpdateListInput

	userId, err := getUsersId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.UpdateList(userId, id, input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUsersId(c)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
 
	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}
