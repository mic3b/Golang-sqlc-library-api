package app

import (
	"example.com/m/internal/services"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/authors", func(c *gin.Context) {
		service := services.NewAuthorService()
		authors, err := service.GetAllAuthors(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, authors)
	})

	r.Run(":8080")
}
