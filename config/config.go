package config

type Configuration struct {
	App        App        `mapstructure:"app" json:"app" yaml:"app"`
	Api        Api        `mapstructure:"api" json:"api" yaml:"api"`
	Log        Log        `mapstructure:"log" json:"log" yaml:"log"`
	Database   string     `mapstructure:"database" json:"database" yaml:"database"`
	MysqlDB    MysqlDB    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	PostgresDB PostgresDB `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
	Jwt        Jwt        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Etcd       Etcd       `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
}
