package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/usecase"
)

// SetRouterGroup defines all the routes for the repo functions
func SetRouterGroup(base *gin.RouterGroup, u usecase.UseCase, working chan bool) *gin.RouterGroup {
	gitGroup := base.Group("/repo")
	{
		gitGroup.POST("/scan", scan(u, working))
		gitGroup.POST("/view", view(u))
	}

	return gitGroup
}
