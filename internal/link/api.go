package link

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"context"
)

func RegisterHandlers(r *gin.Engine, db *sql.DB, ctx context.Context) {
	linkRepo := NewRepository(db)
	linkService := NewService(linkRepo)

	linkRoutes := r.Group("/links")
	{
		linkRoutes.GET("/", func(c *gin.Context) {
			links, err := linkService.All(ctx)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error!",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": links,
			})
		})
		linkRoutes.POST("/", func(c *gin.Context) {
			var createLinkDto CreateDto
			if err := c.ShouldBindJSON(&createLinkDto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			link, err := linkService.Create(ctx, createLinkDto)
			if err != nil {
				c.JSON(500, gin.H{
					"message": "Error!",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": link,
			})
		})
	}
}

func RegisterCatchAll(r *gin.Engine, db *sql.DB, ctx context.Context) {
	linkRepo := NewRepository(db)
	linkService := NewService(linkRepo)

	r.GET("/l/:code", func(c *gin.Context) {
		linkCode := c.Param("code")

		link, err := linkService.GetLinkDestination(ctx, linkCode)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Link code not found!",
			})
			return
		}

		c.Redirect(http.StatusPermanentRedirect, link)
	})
}
