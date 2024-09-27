package config

type App struct {
	Env     EnvConfig `mapstructure:"env" json:"env" yaml:"env"`
	Port    int       `mapstructure:"port" json:"port" yaml:"port"`
	AppName string    `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string    `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

type EnvConfig struct {
	Name  string `mapstructure:"name" json:"name" yaml:"name"`
	Debug bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
}
