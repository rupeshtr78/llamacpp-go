package llamago

import (
	"github.com/rs/zerolog/log"

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

func LlamaRun(serverName string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Error reading config file, %s", err)
	}

	config := ServerConfig{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Msgf("Unable to decode into struct, %v", err)
	}

	server := config.LLmServers

	if len(server) == 0 {
		log.Fatal().Msg("No servers found in config")
	}

	var m ModelConfig
	for _, s := range server {
		if s.Name == serverName {
			m = s
			break
		}
	}

	context := &ModelContext{}
	context.SetStrategy(&LlamaModelStrategy{ModelConfig: m})
	err = context.ExecuteStrategy()
	if err != nil {
		log.Fatal().Msgf("Error executing strategy, %v", err)
	}
}
