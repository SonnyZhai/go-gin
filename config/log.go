package config

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`                   // 日志级别
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`          // 日志文件存放目录
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`          // 日志文件名称
	Format     string `mapstructure:"format" json:"format" yaml:"format"`                // 日志格式
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`       // 是否显示行号
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"` // 最大备份数
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`          // 单个文件最大大小MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`             // 文件最大保存时间
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`          // 是否压缩
}
