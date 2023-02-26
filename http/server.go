package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitscan/config"
	"github.com/gitscan/http/git"
	"github.com/gitscan/internal/database"
	"github.com/gitscan/internal/middleware"
	"github.com/gitscan/internal/usecase"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer() {
	s := run()
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Msg("Server forced to shutdown")
	}
	log.Info().Msg("Server exiting")
}

func run() (s *http.Server) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	db, err := database.New(cfg.DB)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	port := fmt.Sprintf(":%d", cfg.App.Port)

	ginEngine := gin.New()
	ginEngine.Use(middleware.NewLogger())
	ginEngine.Use(gin.Recovery())

	// maximum process if there are more request than MaxProcess will be waiting until processing request finish
	workingChan := make(chan bool, cfg.App.MaxProcess)
	v1 := ginEngine.Group("/v1")

	u := usecase.New(db)
	git.SetRouterGroup(v1, u, workingChan)

	s = &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        ginEngine,
	}

	err = Recovery(u, workingChan)
	if err != nil {
		log.Warn().Msgf("recovery process has an error: %s", err.Error())
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s", err)
		}
	}()

	return
}
