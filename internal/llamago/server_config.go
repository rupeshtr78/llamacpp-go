package llamago

import (
	"reflect"
)

type ServerConfig struct {
	LLmServers []ModelConfig `mapstructure:"servers"`
}

type ModelConfig struct {
	Name      string `mapstructure:"name"`
	Model     string `mapstructure:"model"`
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Threads   string `mapstructure:"threads"`
	CtxSize   string `mapstructure:"ctx_size"`
	BatchSize string `mapstructure:"batch_size"`
	Embedding bool   `mapstructure:"embedding"`
	APIKey    string `mapstructure:"api_key"`
}

func (m ModelConfig) GetArguments() []string {
	args := []string{}
	argValues := reflect.ValueOf(m)
	argTypes := reflect.TypeOf(m)

	for i := 0; i < argValues.NumField(); i++ {
		field := argTypes.Field(i)
		tag := field.Tag.Get("mapstructure")

		if tag == "" {
			continue
		}

		argName := "--" + tag
		argValue := argValues.Field(i)

		// if m.Embedding is false, skip adding the --embedding flag to the args
		if argName == "--embedding" && !argValue.Bool() {
			continue
		} else if argName == "--embedding" && argValue.Bool() {
			args = append(args, argName)
			continue
		}

		if argValue.String() != "" {
			args = append(args, argName, argValue.String())
		}

	}
	return args
}
