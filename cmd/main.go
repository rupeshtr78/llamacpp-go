package main

import (
	"fmt"
	"llama-go/internal/llamago"
)

func main() {
	fmt.Println("Hello, llama-go")
	llamago.LlamaRun("llm_server")
}
