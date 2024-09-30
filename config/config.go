package config

type Configuration struct {
	App        App        `mapstructure:"app" json:"app" yaml:"app"`
	Log        Log        `mapstructure:"log" json:"log" yaml:"log"`
	Database   string     `mapstructure:"database" json:"database" yaml:"database"`
	MysqlDB    MysqlDB    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgresDB PostgresDB `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
}
