package llamago

import (
	"bufio"
	"fmt"
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
	log.Info().Msgf("Running executable %s", constants.LlamaCppCli)
	log.Debug().Msgf("Command: %v", cmd.String())

	// Get the output pipe
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get stdout pipe for %s", constants.LlamaCppServer)
		return err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Error().Err(err).Msgf("Failed to start %s", constants.LlamaCppServer)
		return err
	}

	// Stream the output in real-time
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		output := scanner.Text()
		fmt.Println(output)
	}

	// Check for errors while reading the output
	if err := scanner.Err(); err != nil {
		log.Error().Err(err).Msg("Error reading command output")
		return err
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Error().Err(err).Msgf("Command %s did not finish successfully", constants.LlamaCppServer)
		return err
	}

	return nil
}
