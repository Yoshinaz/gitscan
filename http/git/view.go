package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/database"
	repo2 "github.com/gitscan/internal/service/repo"
	"net/http"
)

func view(r repo2.Interface, db database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req repo2.Request

		// convert request body to request object
		if err := c.BindJSON(&req); err != nil {
			return
		}

		rule := getRules(req.RulesID)

		r.Init(req.Name, req.URL, rule)
		report, err := r.ViewReport(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, report)

			return
		}

		c.JSON(http.StatusOK, report)
	}
}
