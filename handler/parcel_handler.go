package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterParcelRoutes(r *gin.Engine) {
	grp := r.Group("/parcels")
	{
		grp.GET("", rootHandler)
	}

	userGroup := r.Group("/users")
	{
		userGroup.GET("", func(c *gin.Context) {
			c.String(200, "Hello %d", 200)
		})
	}
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"greetings": []string{"hello", "world"},
		"items":     []string{"item1", "item2"},
	})
}
