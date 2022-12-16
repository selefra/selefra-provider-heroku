package heroku_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	Email  string `yaml:"email"  mapstructure:"email"`
	APIKey string `yaml:"api_key"  mapstructure:"api_key"`
}
