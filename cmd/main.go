package main

import (
	"fmt"
	"llama-go/internal/llamago"
	"sync"
)

func main() {
	fmt.Println("Hello, llama-go")

	llamaServers := []string{}
	llamaServers = append(llamaServers, "llm_server_v2")
	// llamaServers = append(llamaServers, "embedding_server")

	var wg sync.WaitGroup
	errChan := make(chan error, len(llamaServers))

	for _, serverName := range llamaServers {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			errChan <- llamago.LlamaRun(name)
		}(serverName)

	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			fmt.Println(err)
		}
	}

}
