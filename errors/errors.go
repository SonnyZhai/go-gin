package errors

import (
	"go-gin/app/common/response"
	"go-gin/cons"
	"go-gin/global"

	"github.com/gin-gonic/gin"
)

// 预定义错误
var (
	// 服务级错误
	ErrServerInternalServerError = global.NewCustomError(500, cons.ERROR_CODE_SERVER_INTERNAL_SERVER_ERROR, "服务内部错误", nil)
	ErrServerDBConnection        = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_CONNECTION, "数据库连接错误", nil)
	ErrServerDBQuery             = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_QUERY, "数据库查询错误", nil)
	ErrServerDBInsert            = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_INSERT, "数据库插入错误", nil)
	ErrServerDBUpdate            = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_UPDATE, "数据库更新错误", nil)
	ErrServerDBDelete            = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_DELETE, "数据库删除错误", nil)
	ErrServerDBTransaction       = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_TRANSACTION, "数据库事务错误", nil)
	ErrServerDBDuplicate         = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_DUPLICATE, "数据库重复错误", nil)
	ErrServerDBDeadlock          = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_DEADLOCK, "数据库死锁错误", nil)
	ErrServerDBUnknown           = global.NewCustomError(500, cons.ERROR_CODE_SERVER_DB_UNKNOWN, "数据库未知错误", nil)
	ErrServerRedisConnection     = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_CONNECTION, "Redis 连接错误", nil)
	ErrServerRedisQuery          = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_QUERY, "Redis 查询错误", nil)
	ErrServerRedisInsert         = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_INSERT, "Redis 插入错误", nil)
	ErrServerRedisUpdate         = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_UPDATE, "Redis 更新错误", nil)
	ErrServerRedisDelete         = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_DELETE, "Redis 删除错误", nil)
	ErrServerRedisTransaction    = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_TRANSACTION, "Redis 事务错误", nil)
	ErrServerRedisDuplicate      = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_DUPLICATE, "Redis 重复错误", nil)
	ErrServerRedisUnknown        = global.NewCustomError(500, cons.ERROR_CODE_SERVER_REDIS_UNKNOWN, "Redis 未知错误", nil)
	ErrServerMQConnection        = global.NewCustomError(500, cons.ERROR_CODE_SERVER_MQ_CONNECTION, "消息队列连接错误", nil)
	ErrServerMQUnknown           = global.NewCustomError(500, cons.ERROR_CODE_SERVER_MQ_UNKNOWN, "消息队列未知错误", nil)
	ErrServerRPCConnection       = global.NewCustomError(500, cons.ERROR_CODE_SERVER_RPC_CONNECTION, "RPC 连接错误", nil)
	ErrServerNotSupported        = global.NewCustomError(501, cons.ERROR_CODE_SERVER_NOT_SUPPORTED, "不支持的服务", nil)
	ErrServerBadGateway          = global.NewCustomError(502, cons.ERROR_CODE_SERVER_BAD_GATEWAY, "网关错误", nil)
	ErrServerServiceUnavailable  = global.NewCustomError(503, cons.ERROR_CODE_SERVER_SERVICE_UNAVAILABLE, "服务不可用", nil)
	ErrServerServiceTimeout      = global.NewCustomError(504, cons.ERROR_CODE_SERVER_SERVICE_TIMEOUT, "服务超时", nil)
	ErrServerGatewayTimeout      = global.NewCustomError(504, cons.ERROR_CODE_SERVER_GATEWAY_TIMEOUT, "网关超时", nil)

	ErrServerUnknown = global.NewCustomError(500, cons.ERROR_CODE_SERVER_UNKNOWN, "未知错误", nil)
)

// HandleError 处理错误并返回响应
func HandleError(c *gin.Context, httpStatus int, code cons.ErrorCode, message string, err error) {
	customErr := global.NewCustomError(httpStatus, code, message, err)
	response.Fail(c, customErr)
}

// HandleErrorWithContext 处理错误并返回响应，带上下文信息
func HandleErrorWithContext(c *gin.Context, httpStatus int, code cons.ErrorCode, message string, err error, context map[string]interface{}) {
	customErr := global.NewCustomError(httpStatus, code, message, err)
	for key, value := range context {
		customErr.WithContext(key, value)
	}
	response.Fail(c, customErr)
}
