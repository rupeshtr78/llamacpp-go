package llamago

import (
	"errors"
	"slices"

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

func LlamaRun(serverName string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("Error reading config file, %s", err)
		return err
	}

	config := ServerConfig{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Msgf("Unable to decode into struct, %v", err)
		return err
	}

	server := config.LLmServers

	if len(server) == 0 {
		log.Fatal().Msg("No servers found in config")
		return errors.New("no servers found in config")
	}

	serverNames := []string{}
	for _, s := range server {
		serverNames = append(serverNames, s.Name)
	}

	log.Debug().Msgf("Server names: %v", serverNames)

	if !slices.Contains(serverNames, serverName) {
		log.Fatal().Msgf("Server %s not found in config", serverName)
		return errors.New("server not found in config")
	}

	var lsc ModelConfig
	for _, s := range server {
		if s.Name == serverName {
			lsc = s
			break
		}
	}

	context := &ModelContext{}
	context.SetStrategy(&LlamaModelStrategy{ModelConfig: lsc})
	err = context.ExecuteStrategy()
	if err != nil {
		log.Fatal().Msgf("Error executing strategy, %v", err)
		return err
	}

	return nil
}
