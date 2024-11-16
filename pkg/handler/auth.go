package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qazaqpyn/api-notz/internal/model"
	"github.com/qazaqpyn/api-notz/internal/tools"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	if err := h.services.CreateUser(c, input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "user registered successfully",
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input model.LoginRequest

	if err := c.BindJSON(&input); err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	token, refreshToken, err := h.services.Login(c, input)
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}
	response, err := json.Marshal(map[string]string{
		"token": token,
	})

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (h *Handler) refreshTokens(c *gin.Context) {
	cookie, err := c.Cookie("refresh-token")
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	token, refreshToken, err := h.services.RefreshTokens(c, cookie)
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": token,
	})

	c.Header("Set-Cookie", fmt.Sprintf("refresh-token=%s; HttpOnly", refreshToken))
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}
