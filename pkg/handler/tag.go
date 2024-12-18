package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/internal/tools"
)

func (h *Handler) getAllTags(c *gin.Context) {
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
	userId, ok := c.Get(userCtx)
	if !ok {
		tools.UnAuthorizedHandler(c.Writer)
		return
	}

	var names model.TagsInput
	if err := c.BindJSON(&names); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	input := make([]model.TagInput, len(names.Names))
	// validate input fields
	for idx, value := range names.Names {
		input[idx] = model.TagInput{
			Name: value,
		}
		if err := input[idx].Validate(); err != nil {
			tools.RequestErrorHandler(c.Writer, err)
			return
		}
	}

	err := h.services.CreateTags(c, userId.(string), input)
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "tags created successfully",
	})
}

func (h *Handler) getUserTags(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		tools.UnAuthorizedHandler(c.Writer)
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
	userId, ok := c.Get(userCtx)
	if !ok {
		tools.UnAuthorizedHandler(c.Writer)
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

	if err := h.services.UpdateTag(c, userId.(string), tagId, &input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "tag updated successfully",
	})
}

func (h *Handler) deleteTag(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		tools.UnAuthorizedHandler(c.Writer)
		return
	}

	tagId := c.Param("id")
	if err := h.services.DeleteTag(c, userId.(string), tagId); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "tag deleted successfully",
	})
}
