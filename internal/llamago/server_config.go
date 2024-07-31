package llamago

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
