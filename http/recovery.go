package http

import (
	"github.com/gitscan/internal/usecase"
	"github.com/rs/zerolog/log"
)

func Recovery(u usecase.UseCase, working chan bool) error {
	log.Info().Msg("recovery process begin")
	infos, err := u.DB.Info().FindRecoveryInfo()
	if err != nil {
		return err
	}

	log.Info().Msgf("Number of recovery process: %d", len(infos))
	for _, info := range infos {
		log.Info().Msgf("recovery for: %s", info.URL)
		u.Recovery(info, working)
	}

	return err
}
