package config

type Api struct {
	Prefix  string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
}
