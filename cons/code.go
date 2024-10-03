package cons

// 错误码规则:
// 1. 错误码为 5 位数:
//              -----------------------------------------------------------------------------
//                  第1位               2、3位                  4、5位
//              -----------------------------------------------------------------------------
//			  服务类型(1 为服务级错误)    通用服务模块代码(00)              具体错误代码
//			  服务类型(1 为服务级错误)    服务用户模块代码(01)              具体错误代码
//        	  业务类型(2 为业务级错误)    通用业务模块代码(00)              具体错误代码
//        	  业务类型(2 为业务级错误)    业务用户模块代码(01)              具体错误代码
//              -----------------------------------------------------------------------------

// ErrorCode 定义错误码类型
type ErrorCode int

// 通用服务模块代码(100)
const (
	// 服务级错误码
	ERROR_CODE_SERVER_INTERNAL_SERVER_ERROR ErrorCode = 10000 + iota // 服务内部错误
	ERROR_CODE_SERVER_SERVICE_UNAVAILABLE                            // 服务不可用
	ERROR_CODE_SERVER_SERVICE_TIMEOUT                                // 服务超时
	ERROR_CODE_SERVER_BAD_GATEWAY                                    // 网关错误
	ERROR_CODE_SERVER_GATEWAY_TIMEOUT                                // 网关超时
	ERROR_CODE_SERVER_NOT_SUPPORTED                                  // 不支持的服务
	ERROR_CODE_SERVER_DB_CONNECTION                                  // 数据库连接错误
	ERROR_CODE_SERVER_DB_QUERY                                       // 数据库查询错误
	ERROR_CODE_SERVER_DB_INSERT                                      // 数据库插入错误
	ERROR_CODE_SERVER_DB_UPDATE                                      // 数据库更新错误
	ERROR_CODE_SERVER_DB_DELETE                                      // 数据库删除错误
	ERROR_CODE_SERVER_DB_TRANSACTION                                 // 数据库事务错误
	ERROR_CODE_SERVER_DB_DUPLICATE                                   // 数据库重复错误
	ERROR_CODE_SERVER_DB_DEADLOCK                                    // 数据库死锁错误
	ERROR_CODE_SERVER_DB_UNKNOWN                                     // 数据库未知错误
	ERROR_CODE_SERVER_REDIS_CONNECTION                               // Redis 连接错误
	ERROR_CODE_SERVER_REDIS_QUERY                                    // Redis 查询错误
	ERROR_CODE_SERVER_REDIS_INSERT                                   // Redis 插入错误
	ERROR_CODE_SERVER_REDIS_UPDATE                                   // Redis 更新错误
	ERROR_CODE_SERVER_REDIS_DELETE                                   // Redis 删除错误
	ERROR_CODE_SERVER_REDIS_TRANSACTION                              // Redis 事务错误
	ERROR_CODE_SERVER_REDIS_DUPLICATE                                // Redis 重复错误
	ERROR_CODE_SERVER_REDIS_UNKNOWN                                  // Redis 未知错误
	ERROR_CODE_SERVER_MQ_CONNECTION                                  // 消息队列连接错误
	ERROR_CODE_SERVER_MQ_UNKNOWN                                     // 消息队列未知错误
	ERROR_CODE_SERVER_RPC_CONNECTION                                 // RPC 连接错误

	ERROR_CODE_SERVER_UNKNOWN ErrorCode = 10099
)

// 服务用户模块代码(101)
const (
	// 用户模块服务级错误码
	ERROR_CODE_SERVER_USER_UNAUTHORIZED      ErrorCode = 10100 + iota // 用户未认证
	ERROR_CODE_SERVER_USER_FORBIDDEN                                  // 用户未授权
	ERROR_CODE_SERVER_USER_INVALID_TOKEN                              // 用户令牌无效
	ERROR_CODE_SERVER_USER_FAILED_TOKEN                               // 创建用户令牌失败
	ERROR_CODE_SERVER_USER_INVALID_SESSION                            // 用户会话无效
	ERROR_CODE_SERVER_USER_INVALID_COOKIE                             // 用户 Cookie 无效
	ERROR_CODE_SERVER_REGISTER_FAILED                                 // 用户注册失败
	ERROR_CODE_SERVER_LOGIN_FAILED                                    // 用户登录失败
	ERROR_CODE_SERVER_USER_OPERATION_FAILED                           // 用户操作失败
	ERROR_CODE_SERVER_USER_TOO_MANY_REQUESTS                          // 用户请求过多

	ERROR_CODE_SERVER_USER_UNKNOWN ErrorCode = 10199
)

// 通用业务模块代码(200)
const (
	ERROR_CODE_BUSINESS_INVALID_PARAMS   ErrorCode = 20000 + iota // 无效参数
	ERROR_CODE_BUSINESS_INVALID_METHOD                            // 无效方法
	ERROR_CODE_BUSINESS_INVALID_TOKEN                             // 无效令牌
	ERROR_CODE_BUSINESS_INVALID_GUARD                             // 无效守卫
	ERROR_CODE_BUSINESS_INVALID_SESSION                           // 无效会话
	ERROR_CODE_BUSINESS_INVALID_URL                               // 无效 URL
	ERROR_CODE_BUSINESS_INVALID_FILE                              // 无效文件
	ERROR_CODE_BUSINESS_INVALID_IMAGE                             // 无效图片
	ERROR_CODE_BUSINESS_INVALID_VIDEO                             // 无效视频
	ERROR_CODE_BUSINESS_INVALID_DOCUMENT                          // 无效文档
	ERROR_CODE_BUSINESS_INVALID_CHARSET                           // 无效字符集
	ERROR_CODE_BUSINESS_OPERATION

	ERROR_CODE_BUSINESS_UNKNOWN ErrorCode = 20099
)

// 业务用户模块代码(201)
const (
	ERROR_CODE_BUSINESS_ACCOUNT_NOT_FOUND          ErrorCode = 20100 + iota // 账户未找到
	ERROR_CODE_BUSINESS_ACCOUNT_LOCKED                                      // 用户账户被锁定
	ERROR_CODE_BUSINESS_ACCOUNT_DISABLED                                    // 用户账户被禁用
	ERROR_CODE_BUSINESS_ACCOUNT_ALREADY_EXISTS                              // 用户账户已存在
	ERROR_CODE_BUSINESS_ACCOUNT_PASSWORD_INCORRECT                          // 密码不正确
	ERROR_CODE_BUSINESS_ACCOUNT_MOBILE_INCORRECT                            // 手机号不正确
	ERROR_CODE_BUSINESS_ACCOUNT_EMAIL_INCORRECT                             // 邮箱不正确
	ERROR_CODE_BUSINESS_ACCOUNT_CAPTCHA_INCORRECT                           // 验证码不正确

	ERROR_CODE_BUSINESS_ACCOUNT_UNKNOWN ErrorCode = 20199
)
