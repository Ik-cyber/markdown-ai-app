package v1

import (
	"github.com/Ik-cyber/markdown-ai-app/internal/markdown"
	"github.com/Ik-cyber/markdown-ai-app/internal/middleware"
	"github.com/Ik-cyber/markdown-ai-app/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", user.Register)
	rg.POST("/login", user.Login)

	auth := rg.Group("")
	auth.Use(middleware.JWTAuthMiddleware())

	auth.POST("/markdowns", markdown.CreateMarkdown)
	auth.GET("/markdowns", markdown.ListMarkdowns)
	auth.GET("/markdowns/:id", markdown.GetMarkdown)
	auth.PUT("/markdowns/:id", markdown.UpdateMarkdown)
	auth.DELETE("/markdowns/:id", markdown.DeleteMarkdown)
}
