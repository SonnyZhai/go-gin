package config

type Configuration struct {
	App        App        `mapstructure:"app" json:"app" yaml:"app"`
	Log        Log        `mapstructure:"log" json:"log" yaml:"log"`
	MysqlDB    MysqlDB    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgresDB PostgresDB `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
}
