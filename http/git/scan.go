package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/service/repo"
	"net/http"
)

func scan(db database.DB, working chan bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req repo.Request

		// convert request body to request object
		if err := c.BindJSON(&req); err != nil {
			return
		}

		rule := GetRules(req.RulesID)

		r := repo.New(req.Name, req.URL, rule)
		status, err := repo.Scan(r, db, working)
		if err != nil {
			c.JSON(http.StatusInternalServerError, status)

			return
		}

		c.JSON(http.StatusOK, status)
	}
}
