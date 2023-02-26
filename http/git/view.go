package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/usecase"
	"net/http"
)

func view(u usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request

		// convert request body to request object
		if err := c.BindJSON(&req); err != nil {
			return
		}

		report, err := u.View(req.Name, req.URL, req.AllCommit, req.RulesSet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, report)

			return
		}

		c.JSON(http.StatusOK, report)
	}
}
