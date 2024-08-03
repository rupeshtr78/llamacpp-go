package main

import (
	"fmt"
	"llama-go/internal/llamago"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello, llama-go")

	errChan := make(chan error)
	go func(name string) {
		errChan <- llamago.LlamaRun(name)
	}("llm_server")

	err := <-errChan
	if err != nil {
		log.Fatal().Msgf("Error running llama-go, %v", err)
	}

}
