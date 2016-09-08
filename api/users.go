package api

import (
	"github.com/gin-gonic/gin"

	"github.com/AusDTO/lgtm/router/middleware/session"
)

// GetUser gets the currently authenticated user.
func GetUser(c *gin.Context) {
	c.JSON(200, session.User(c))
}
