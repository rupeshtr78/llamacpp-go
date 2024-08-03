package llamago

import (
	"llama-go/internal/constants"
	"os/exec"

	"github.com/rs/zerolog/log"
)

type ModelStrategy interface {
	Execute() error
}

type LlamaModelStrategy struct {
	ModelConfig ModelConfig
}

func (l *LlamaModelStrategy) Execute() error {
	args := l.ModelConfig.GetArguments()
	log.Info().Msgf("Executing llama-cli with args: %v", args)

	cmd := exec.Command(constants.LlamaCppCli, args...)
	err := cmd.Run()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to run %s", constants.LlamaCppServer)

	}
	return err
}
