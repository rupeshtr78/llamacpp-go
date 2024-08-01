package llamago

import (
	"log"

	"github.com/spf13/viper"
)

type ModelContext struct {
	strategy ModelStrategy
}

func (mc *ModelContext) SetStrategy(strategy ModelStrategy) {
	mc.strategy = strategy
}

func (mc *ModelContext) ExecuteStrategy() error {
	return mc.strategy.Execute()
}

func LlamaStrategy() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

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

	var m ModelConfig
	for _, s := range server {
		if s.Name == "llm_server" {
			m = s
			break
		}
	}

	context := &ModelContext{}
	context.SetStrategy(&LlamaModelStrategy{ModelConfig: m})
	err = context.ExecuteStrategy()
	if err != nil {
		log.Fatalf("Failed to execute strategy: %v", err)
	}
}
