package cons

// 定义一些全局常量
const (
	// 项目初始化配置
	CONFIG_PATH = "CONFIG_PATH" // 环境变量的配置文件路径
	APP_ENV     = "APP_ENV"     // 环境变量的运行环境
	ENV_PROD    = "prod"        // 生产环境
	CONFIG_PROD = "config.prod" // 生产环境配置文件名称
	ENV_TEST    = "test"        // 测试环境
	CONFIG_TEST = "config.test" // 测试环境配置文件名称
	CONFIG_DEV  = "settings"    // 开发环境配置文件名称
	YAML_TYPE   = "yaml"        // YAML 配置文件类型
	TOML_TYPE   = "toml"        // TOML 配置文件类型
	JSON_TYPE   = "json"        // JSON 配置文件类型

	// 符号名称
	COLON                = ":"  // 冒号
	LEFT_ROUND_BRACKET   = "("  // 左括号
	RIGHT_ROUND_BRACKET  = ")"  // 右括号
	LEFT_SQUARE_BRACKET  = "["  // 左方括号
	RIGHT_SQUARE_BRACKET = "]"  // 右方括号
	LEFT_BRACE           = "{"  // 左大括号
	RIGHT_BRACE          = "}"  // 右大括号
	QUESTION_MARK        = "?"  // 问号
	COMMA                = ","  // 逗号
	SEMICOLON            = ";"  // 分号
	PERCENT              = "%"  // 百分号
	ASTERISK             = "*"  // 星号
	EXCLAMATION          = "!"  // 感叹号
	SLASH                = "/"  // 斜杠
	DOT                  = "."  // 点
	EQUAL                = "="  // 等号
	AND                  = "&"  // 与
	STRING_PLACEHOLDER   = "%s" // 字符串占位符

	// 日期时间格式
	DateTimeFormat = "2006-01-02 15:04:05.000"
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05.000"

	// 日志级别
	LogLevelDebug  = "debug"
	LogLevelInfo   = "info"
	LogLevelWarn   = "warn"
	LogLevelError  = "error"
	LogLevelDPanic = "dpanic"
	LogLevelPanic  = "panic"
	LogLevelFatal  = "fatal"
	LogLevelSilent = "silent"

	// 数据库相关
	DATABASE_AT_TCP                  = "@tcp"
	DATABASE_CHARSET                 = "charset"
	DATABASE_TYPE_MYSQL              = "mysql"
	DATABASE_TYPE_POSTGRESQL         = "postgres"
	DATABASE_POSTGRESQL_TABLE_PREFIX = "t_"
	DATABASE_MYSQL_PARAMS            = "&parseTime=true&loc=Local"

	// 致命信息输出相关
	FATAL_READ_CONFIG      = "读取配置文件失败: "
	FATAL_CONFIG_TO_GLOBAL = "配置赋值给全局变量失败: "

	// 错误信息输出相关
	ERROR_MYSQL_DB_CONNECT    = "Mysql数据库连接失败: "
	ERROR_POSTGRES_DB_CONNECT = "Postgres数据库连接失败: "
	ERROR_DB_CONFIG_DBNAME    = "数据库名称为空，连接失败"
	ERROR_READ_CONFIG         = "读取配置文件失败: "
	ERROR_RELOAD_CONFIG       = "重新加载配置文件失败: "
	ERROR_DB_TYPE_UNSUPPORT   = "不支持的数据库类型: "

	// 信息打印相关
	INFO_READ_CONFIG      = "读取配置文件成功: "
	INFO_MODIFY_CONFIG    = "配置文件已修改并重新加载: "
	INFO_RELOAD_CONFIG    = "重新加载配置文件成功: "
	INFO_DB_CONNECT       = "数据库连接成功"
	INFO_LOG_INIT_SUCCESS = "日志初始化成功"
)
