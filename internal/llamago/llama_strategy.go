package llamago

import (
	"log"
	"os/exec"
)

type ModelStrategy interface {
	Execute() error
}

type LlamaModelStrategy struct {
	ModelConfig ModelConfig
}

func (l *LlamaModelStrategy) Execute() error {
	args := []string{
		"--model", l.ModelConfig.Model,
		"--host", l.ModelConfig.Host,
		"--port", l.ModelConfig.Port,
		"--threads", l.ModelConfig.Threads,
		"--ctx-size", l.ModelConfig.CtxSize,
		"--batch-size", l.ModelConfig.BatchSize,
	}

	if l.ModelConfig.Embedding {
		args = append(args, "--embedding")
	}

	args = append(args, "--api-key", l.ModelConfig.APIKey)

	cmd := exec.Command("llama-cli", args...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	return err
}
