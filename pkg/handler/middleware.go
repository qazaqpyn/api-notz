package handler

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qazaqpyn/api-notz/internal/tools"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

var UnAuthorizedError = errors.New("unauthorized")

// Authorizarion is a middleware that checks if the user is authorized
func (h *Handler) authorizarion(c *gin.Context) {
	authHeader := c.GetHeader(authorizationHeader)
	if authHeader == "" {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}

	if len(headerParts[1]) == 0 {
		tools.RequestErrorHandler(c.Writer, UnAuthorizedError)
		return
	}
	userId, err := h.services.ParseToken(c, headerParts[1])
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		return
	}

	c.Set(userCtx, userId)
	c.Next()
}
