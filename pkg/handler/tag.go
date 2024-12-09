package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/internal/tools"
)

func (h *Handler) getAllTags(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}

	tags, err := h.services.GetAllTags(c)
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": &tags,
	})
}

func (h *Handler) createTags(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}

	var input []model.TagInput
	if err := c.BindJSON(&input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	// validate input fields
	for _, value := range input {
		if err := value.Validate(); err != nil {
			tools.RequestErrorHandler(c.Writer, err)
			return
		}
	}

	tags, err := h.services.CreateTags(c, &input)
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"data": tags,
	})
}

func (h *Handler) getUserTags(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}

	tags, err := h.services.GetUserTags(c, userId.(string))
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": tags,
	})
}

func (h *Handler) updateTag(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}

	tagId := c.Param("id")
	var input model.TagInput
	if err := c.BindJSON(&input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	// validate
	if err := input.Validate(); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	if err := h.services.UpdateTag(c, tagId, &input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "tag updated successfully",
	})
}

func (h *Handler) deleteTag(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}

	tagId := c.Param("id")
	if err := h.services.DeleteTag(c, tagId); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "tag deleted successfully",
	})
}
