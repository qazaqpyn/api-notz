package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qazaqpyn/api-notz/internal/tools"
	"github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

// Authorizarion is a middleware that checks if the user is authorized
func (h *Handler) authorizarion(c *gin.Context) {
	authHeader := c.GetHeader(authorizationHeader)
	if authHeader == "" {
		tools.UnAuthorizedHandler(c.Writer)
		c.Abort()
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		tools.UnAuthorizedHandler(c.Writer)
		c.Abort()
		return
	}

	if len(headerParts[1]) == 0 {
		tools.UnAuthorizedHandler(c.Writer)
		c.Abort()
		return
	}
	userId, err := h.services.ParseToken(c, headerParts[1])
	if err != nil {
		tools.RequestErrorHandler(c.Writer, err)
		c.Abort()
		return
	}

	c.Set(userCtx, userId)
	c.Next()
}

func (h *Handler) printRequest(c *gin.Context) {
	logrus.Printf(c.FullPath())
	c.Next()
}
