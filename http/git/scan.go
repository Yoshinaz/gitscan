package git

import (
	"github.com/gin-gonic/gin"
	"github.com/gitscan/internal/usecase"
	"github.com/rs/zerolog/log"
	"net/http"
)

func scan(u usecase.UseCase, working chan bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request

		if err := c.BindJSON(&req); err != nil {
			log.Warn().Msgf("request error %s", err.Error())
			return
		}

		status, err := u.Scan(req.Name, req.URL, req.RulesSet, req.AllCommit, working)
		if err != nil {
			c.JSON(http.StatusInternalServerError, status)

			return
		}

		c.JSON(http.StatusOK, map[string]string{"status": status})
	}
}
