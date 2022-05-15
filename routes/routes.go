package routes

import (
	user "gin-gionic-1/controllers/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartService() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		api.GET("/users", user.GetAllUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8080")
}
