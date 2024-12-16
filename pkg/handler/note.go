package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/internal/tools"
)

func (h *Handler) getAllNotes(c *gin.Context) {
	notes, err := h.services.GetAllNotes(c)
	if err != nil {
		tools.UnAuthorizedHandler(c.Writer)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": &notes,
	})
}

func (h *Handler) createNote(c *gin.Context) {
	var input model.Note
	if err := c.BindJSON(&input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	note, err := h.services.CreateNote(c, &input)
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"data": note,
	})
}

func (h *Handler) getNoteById(c *gin.Context) {
	noteId := c.Param("id")
	note, err := h.services.GetNoteById(c, noteId)
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": note,
	})
}

func (h *Handler) updateNote(c *gin.Context) {
	noteId := c.Param("id")
	var input model.UpdateNoteInput
	if err := c.BindJSON(&input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	if err := h.services.UpdateNote(c, noteId, &input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "note updated successfully",
	})
}

func (h *Handler) deleteNote(c *gin.Context) {
	noteId := c.Param("id")
	if err := h.services.DeleteNote(c, noteId); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "note deleted successfully",
	})
}
