package config

type Etcd struct {
	Token           string `mapstructure:"token" json:"token" yaml:"token"`
	AccountId       string `mapstructure:"account_id" json:"account_id" yaml:"account_id"`
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret"`
	DefaultEndpoint string `mapstructure:"default_endpoint" json:"default_endpoint" yaml:"default_endpoint"`
	EuEndpoint      string `mapstructure:"eu_endpoint" json:"eu_endpoint" yaml:"eu_endpoint"`
}
