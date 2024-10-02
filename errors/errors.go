package errors

import (
	"go-gin/cons"
	"go-gin/global"
)

// 预定义错误
var (
	// 服务级错误
	ErrServerInternalServerError = global.NewCustomError(cons.ERROR_CODE_SERVER_INTERNAL_SERVER_ERROR, "服务内部错误", 500, nil)
	ErrServerDBConnection        = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_CONNECTION, "数据库连接错误", 500, nil)
	ErrServerDBQuery             = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_QUERY, "数据库查询错误", 500, nil)
	ErrServerDBInsert            = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_INSERT, "数据库插入错误", 500, nil)
	ErrServerDBUpdate            = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_UPDATE, "数据库更新错误", 500, nil)
	ErrServerDBDelete            = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_DELETE, "数据库删除错误", 500, nil)
	ErrServerDBTransaction       = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_TRANSACTION, "数据库事务错误", 500, nil)
	ErrServerDBDuplicate         = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_DUPLICATE, "数据库重复错误", 500, nil)
	ErrServerDBDeadlock          = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_DEADLOCK, "数据库死锁错误", 500, nil)
	ErrServerDBUnknown           = global.NewCustomError(cons.ERROR_CODE_SERVER_DB_UNKNOWN, "数据库未知错误", 500, nil)
	ErrServerRedisQuery          = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_QUERY, "Redis 查询错误", 500, nil)
	ErrServerRedisInsert         = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_INSERT, "Redis 插入错误", 500, nil)
	ErrServerRedisUpdate         = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_UPDATE, "Redis 更新错误", 500, nil)
	ErrServerRedisDelete         = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_DELETE, "Redis 删除错误", 500, nil)
	ErrServerRedisTransaction    = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_TRANSACTION, "Redis 事务错误", 500, nil)
	ErrServerRedisDuplicate      = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_DUPLICATE, "Redis 重复错误", 500, nil)
	ErrServerRedisUnknown        = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_UNKNOWN, "Redis 未知错误", 500, nil)
	ErrServerMQConnection        = global.NewCustomError(cons.ERROR_CODE_SERVER_MQ_CONNECTION, "消息队列连接错误", 500, nil)
	ErrServerMQUnknown           = global.NewCustomError(cons.ERROR_CODE_SERVER_MQ_UNKNOWN, "消息队列未知错误", 500, nil)
	ErrServerRPCConnection       = global.NewCustomError(cons.ERROR_CODE_SERVER_RPC_CONNECTION, "RPC 连接错误", 500, nil)
	ErrServerServiceUnavailable  = global.NewCustomError(cons.ERROR_CODE_SERVER_SERVICE_UNAVAILABLE, "服务不可用", 503, nil)
	ErrServerServiceTimeout      = global.NewCustomError(cons.ERROR_CODE_SERVER_SERVICE_TIMEOUT, "服务超时", 504, nil)
	ErrServerBadGateway          = global.NewCustomError(cons.ERROR_CODE_SERVER_BAD_GATEWAY, "网关错误", 502, nil)
	ErrServerGatewayTimeout      = global.NewCustomError(cons.ERROR_CODE_SERVER_GATEWAY_TIMEOUT, "网关超时", 504, nil)
	ErrServerNotSupported        = global.NewCustomError(cons.ERROR_CODE_SERVER_NOT_SUPPORTED, "不支持的服务", 501, nil)
	ErrServerRedisConnection     = global.NewCustomError(cons.ERROR_CODE_SERVER_REDIS_CONNECTION, "Redis 连接错误", 500, nil)

	ErrServerUnknown = global.NewCustomError(cons.ERROR_CODE_SERVER_UNKNOWN, "未知错误", 500, nil)

	// 用户模块服务级错误
	ErrServerUserUnauthorized    = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_UNAUTHORIZED, "用户未认证", 401, nil)
	ErrServerUserInvalidToken    = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_INVALID_TOKEN, "用户令牌无效", 401, nil)
	ErrServerUserInvalidSession  = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_INVALID_SESSION, "用户会话无效", 401, nil)
	ErrServerUserInvalidCookie   = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_INVALID_COOKIE, "用户 Cookie 无效", 401, nil)
	ErrServerUserForbidden       = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_FORBIDDEN, "用户未授权", 403, nil)
	ErrServerUserTooManyRequests = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_TOO_MANY_REQUESTS, "用户请求过多", 429, nil)
	ErrServerUserOperationFailed = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_OPERATION_FAILED, "用户操作失败", 500, nil)

	ErrServerUserUnknown = global.NewCustomError(cons.ERROR_CODE_SERVER_USER_UNKNOWN, "未知错误", 500, nil)

	// 业务级错误
	ErrBusinessInvalidParams   = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_PARAMS, "无效参数", 400, nil)
	ErrBusinessInvalidURL      = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_URL, "无效 URL", 400, nil)
	ErrBusinessInvalidFile     = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_FILE, "无效文件", 400, nil)
	ErrBusinessInvalidImage    = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_IMAGE, "无效图片", 400, nil)
	ErrBusinessInvalidVideo    = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_VIDEO, "无效视频", 400, nil)
	ErrBusinessInvalidDocument = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_DOCUMENT, "无效文档", 400, nil)
	ErrBusinessInvalidCharset  = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_CHARSET, "无效字符集", 400, nil)
	ErrBusinessInvalidToken    = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_TOKEN, "无效令牌", 401, nil)
	ErrBusinessInvalidSession  = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_SESSION, "无效会话", 401, nil)
	ErrBusinessInvalidMethod   = global.NewCustomError(cons.ERROR_CODE_BUSINESS_INVALID_METHOD, "无效方法", 405, nil)

	ErrBusinessInvalidUnknown = global.NewCustomError(cons.ERROR_CODE_BUSINESS_UNKNOWN, "未知错误", 400, nil)

	// 用户模块业务级错误
	ErrBusinessAccountNotFound          = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_NOT_FOUND, "账户未找到", 404, nil)
	ErrBusinessAccountLocked            = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_LOCKED, "用户账户被锁定", 403, nil)
	ErrBusinessAccountDisabled          = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_DISABLED, "用户账户被禁用", 403, nil)
	ErrBusinessAccountAlreadyExists     = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_ALREADY_EXISTS, "用户账户已存在", 409, nil)
	ErrBusinessAccountPasswordIncorrect = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_PASSWORD_INCORRECT, "密码不正确", 401, nil)
	ErrBusinessAccountMobileIncorrect   = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_MOBILE_INCORRECT, "手机号不正确", 400, nil)
	ErrBusinessAccountEmailIncorrect    = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_EMAIL_INCORRECT, "邮箱不正确", 400, nil)
	ErrBusinessAccountCaptchaIncorrect  = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_CAPTCHA_INCORRECT, "验证码不正确", 400, nil)

	ErrBusinessAccountUnknown = global.NewCustomError(cons.ERROR_CODE_BUSINESS_ACCOUNT_UNKNOWN, "未知错误", 400, nil)
)
