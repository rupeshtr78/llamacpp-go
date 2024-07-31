package llamago

import (
	"log"
	"os/exec"

	"github.com/spf13/viper"
)

func LlamaMain() {
	// Set the configuration file name and path
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // YAML format
	viper.AddConfigPath("./config/") // look for config in the config directory

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config := ServerConfig{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	server := config.LLmServers

	if len(server) == 0 {
		log.Fatalf("No servers found in the configuration file")
	}

	// Get the first server based on name
	var m ModelConfig
	for _, s := range server {
		if s.Name == "llm_server" {
			m = s
			return
		}
	}

	// Define the command and its arguments
	args := []string{
		"--model", m.Model,
		"--host", m.Host,
		"--port", m.Port,
		"--threads", m.Threads,
		"--ctx-size", m.CtxSize,
		"--batch-size", m.BatchSize,
	}

	// Add the embedding flag if enabled
	if m.Embedding {
		args = append(args, "--embedding")
	}

	// Add the API key
	args = append(args, "--api-key", viper.GetString("api_key"))

	// Create the command
	cmd := exec.Command("llama-cli", args...)

	// Run the command
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
}
