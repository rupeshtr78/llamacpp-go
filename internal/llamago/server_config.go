package llamago

import (
	"reflect"
)

type ServerConfig struct {
	LLmServers []ModelConfig `mapstructure:"servers"`
}

type ModelConfig struct {
	Name             string `mapstructure:"name"`
	Model            string `mapstructure:"model"`
	Host             string `mapstructure:"host"`
	Port             string `mapstructure:"port"`
	Threads          string `mapstructure:"threads"`
	CtxSize          string `mapstructure:"ctx_size"`
	BatchSize        string `mapstructure:"batch_size"`
	Embedding        bool   `mapstructure:"embedding"`
	APIKey           string `mapstructure:"api_key"`
	Verbose          bool   `mapstructure:"verbose"`
	Seed             string `mapstructure:"seed"`
	SystemPromptFile string `mapstructure:"system-prompt-file"`
	ChatTemplate     string `mapstructure:"chat-template"`
	Prompt           string `mapstructure:"prompt"`
	Temperature      string `mapstructure:"temp"`
	TopK             string `mapstructure:"top-k"`
	TopP             string `mapstructure:"top-p"`
	RepeatPenalty    string `mapstructure:"repeat_penalty"`
	GPULayers        string `mapstructure:"gpu-layers"`
	MiroStat         string `mapstructure:"mirostat"`
	Stream           string `mapstructure:"stream"`
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

		if argName == "--verbose" && argValue.Bool() {
			args = append(args, argName)
			continue
		}

		if argName == "--name" {
			continue
		}

		prompt := m.Prompt
		// add double quotes aroung the prompt string
		promptStr := `"` + prompt + `"`
		if m.Prompt != "" {
			args = append(args, argName, promptStr)
			continue
		}

		// convert the value to a string
		argValueStr := argValue.String()

		if argValueStr != "" {
			args = append(args, argName, argValueStr)
		}

	}
	return args
}
