package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/service/repo"
)

// SetRouterGroup defines all the routes for the repo functions
func SetRouterGroup(base *gin.RouterGroup, r repo.Interface, db database.DB, working chan bool) *gin.RouterGroup {
	gitGroup := base.Group("/repo")
	{
		gitGroup.POST("/scan", scan(r, db, working))
		gitGroup.POST("/view", view(r, db))
	}

	return gitGroup
}
