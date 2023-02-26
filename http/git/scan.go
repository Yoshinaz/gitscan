package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/usecase"
	"net/http"
)

func scan(u usecase.UseCase, working chan bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request

		// convert request body to request object
		if err := c.BindJSON(&req); err != nil {
			return
		}

		status, err := u.Scan(req.Name, req.URL, req.RulesSet, req.ScanAllCommit, working)
		if err != nil {
			c.JSON(http.StatusInternalServerError, status)

			return
		}

		c.JSON(http.StatusOK, status)
	}
}
