package web

import (
	"github.com/AusDTO/lgtm/cache"
	"github.com/AusDTO/lgtm/router/middleware/session"
	"github.com/AusDTO/lgtm/shared/token"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user := session.User(c)

	switch {
	case user == nil:
		c.HTML(200, "brand.html", gin.H{})
	default:
		teams, _ := cache.GetTeams(c, user)
		csrf, _ := token.New(token.CsrfToken, user.Login).Sign(user.Secret)
		c.HTML(200, "index.html", gin.H{"user": user, "csrf": csrf, "teams": teams})
	}
}
