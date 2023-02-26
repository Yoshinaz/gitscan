package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/service/repo"
	"net/http"
)

func view(db database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req repo.Request

		// convert request body to request object
		if err := c.BindJSON(&req); err != nil {
			return
		}

		rule := GetRules(req.RulesID)

		r := repo.New(req.Name, req.URL, rule)
		report, err := repo.ViewReport(r, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, report)

			return
		}

		c.JSON(http.StatusOK, report)
	}
}
