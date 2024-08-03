package main

import (
	"fmt"
	"llama-go/internal/llamago"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello, llama-go")

	serverName := "llm_server_v2"

	// errChan := make(chan error)
	// go func(name string) {
	// 	errChan <- llamago.LlamaRun(name)
	// }(serverName)

	// err := <-errChan
	err := llamago.LlamaRun(serverName)
	if err != nil {
		log.Fatal().Msgf("Error running llama-go, %v", err)
	}

}
