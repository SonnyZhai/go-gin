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

	// API相关
	API_USER_GROUP           = "user" // 用户分组
	API_USER_ID              = "id"
	API_AUTH_GROUP           = "auth" // 需要认证的分组
	API_MESSAGE              = "message"
	API_SUCCESS              = "Success"
	API_FAILED               = "Failed"
	API_ERROR                = "error"
	API_AUTH_NAME            = "Authorization"
	API_TOKEN_NAME           = "token"
	API_REFRESH_TOKEN        = "refresh_token"
	API_REFRESH_TOKEN_EXPIRE = "refresh_token_expire"

	// 符号名称
	COLON                = ":"    // 冒号
	LEFT_ROUND_BRACKET   = "("    // 左括号
	RIGHT_ROUND_BRACKET  = ")"    // 右括号
	LEFT_SQUARE_BRACKET  = "["    // 左方括号
	RIGHT_SQUARE_BRACKET = "]"    // 右方括号
	LEFT_BRACE           = "{"    // 左大括号
	RIGHT_BRACE          = "}"    // 右大括号
	QUESTION_MARK        = "?"    // 问号
	COMMA                = ","    // 逗号
	SEMICOLON            = "; "   // 分号
	PERCENT              = "%"    // 百分号
	ASTERISK             = "*"    // 星号
	EXCLAMATION          = "!"    // 感叹号
	SLASH                = "/"    // 斜杠
	DOT                  = "."    // 点
	EQUAL                = "="    // 等号
	AND                  = "&"    // 与
	STRING_PLACEHOLDER   = "%s"   // 字符串占位符
	STRING_PLACEHOLDER_N = "%s\n" // 字符串占位符（换行）
	NUMBER_PLACEHOLDER   = "%d"   // 数字占位符

	// 日期时间格式
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	TIMEZONE       = "Asia/Shanghai"

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

	// OSS 相关
	OSS_R2_NAME = "r2"

	// ETCD 相关
	ETCD_VERSION     = "etcd3"
	ETCD_ENV_ADDR    = "ETCD_ADDR"
	ETCD_CONFIG_PATH = "/v3/config/storage/r2"

	// redis 相关
	JWT_BLACK_LIST     = "jwt_black_list"
	REFRESH_TOKEN_LOCK = "refresh_token_lock"

	// 致命信息输出相关
	FATAL_SERVER_START          = "服务器启动失败: "
	FATAL_SERVER_SHUTDOWN       = "服务器关闭失败: "
	FATAL_READ_CONFIG           = "读取配置文件失败: "
	FATAL_CONFIG_TO_GLOBAL      = "配置赋值给全局变量失败: "
	FATAL_DB_CONNECT            = "数据库连接失败: "
	FATAL_DECODE_POSTGRES       = "解码PostgreSQL配置失败: "
	FATAL_DECODE_MYSQL          = "解码MySQL配置失败: "
	FATAL_READ_ETCD_CONFIG      = "读取etcd配置失败: "
	FATAL_RELOAD_ETCD_CONFIG    = "重新加载etcd配置失败: "
	FATAL_ETCD_ADDR_PROVIDER    = "etcd配置地址为空: "
	FATAL_ADD_REMOTE_PROVIDER   = "添加远程配置提供者失败: "
	FATAL_READ_REMOTE_CONFIG    = "读取远程配置失败: "
	FATAL_REMOTE_VALUE_TO_CONF  = "远程配置赋值给全局变量失败: "
	FATAL_ETCD_CONFIG_TO_GLOBAL = "etcd配置赋值给全局变量失败: "

	// 错误信息输出相关
	ERROR_MYSQL_DB_CONNECT    = "Mysql数据库连接失败: "
	ERROR_POSTGRES_DB_CONNECT = "Postgres数据库连接失败: "
	ERROR_REDIS_CONNECTION    = "Redis连接失败: "
	ERROR_DB_CONFIG_DBNAME    = "数据库名称为空，连接失败"
	ERROR_READ_CONFIG         = "读取配置文件失败: "
	ERROR_RELOAD_CONFIG       = "重新加载配置文件失败: "
	ERROR_DB_TYPE_UNSUPPORT   = "不支持的数据库类型: "
	ERROR_DB_MIGRATE          = "数据库迁移失败: "
	ERROR_DEFAULT_REQUEST     = "参数错误"

	// 用户错误业务逻辑相关
	ERROR_USERNAME_EXIST   = "用户名已存在"
	ERROR_MOBILE_EXIST     = "手机号已存在"
	ERROR_EMAIL_EXIST      = "邮箱已存在"
	ERROR_INVALID_USER_ID  = "无效的用户ID"
	ERROR_USER_NOT_EXIST   = "用户不存在"
	ERROR_REGISTER_FAILED  = "数据库异常，注册失败"
	ERROR_EMPTY_TOKEN      = "token为空, 请登录后访问"
	ERROR_CLAIMS_TOKEN     = "token是无效的, 请重新登录"
	ERROR_INVALID_ISSUER   = "token是错误的, 请重新登录"
	ERROR_LOGIN_FAILED     = "用户不存在或用户名密码错误"
	ERROR_TOKEN_GUARD_NAME = "token中的守卫名称错误: "

	// 信息打印相关
	INFO_SERVER_START       = "服务器启动成功，端口: "
	INFO_SERVER_IN_SHUTDOWN = "接收到关闭信号，服务器正在关闭..."
	INFO_SERVER_SHUTDOWN    = "服务器已关闭"
	INFO_READ_CONFIG        = "读取配置文件成功: "
	INFO_MODIFY_CONFIG      = "配置文件已修改并重新加载: "
	INFO_RELOAD_CONFIG      = "重新加载配置文件成功: "
	INFO_DB_CONNECT         = "数据库连接成功"
	INFO_LOG_INIT_SUCCESS   = "日志初始化成功"
)
