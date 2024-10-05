package bootstrap

import (
	"fmt"
	"go-gin/app/models"
	"go-gin/cons"
	"go-gin/global"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 数据库初始化
func InitializeDB() *gorm.DB {
	// 定义一个错误变量
	var err error
	var db *gorm.DB

	// 根据配置文件选择数据库类型
	switch global.App.Config.Database {
	case cons.DATABASE_TYPE_MYSQL:
		db, err = initMysqlGorm()
	case cons.DATABASE_TYPE_POSTGRESQL:
		db, err = initPostgresGorm()
	default:
		global.App.Log.Fatal(cons.ERROR_DB_TYPE_UNSUPPORT+cons.STRING_PLACEHOLDER, zap.String("数据库类型", global.App.Config.Database))
	}

	// 如果有错误则输出错误信息
	if err != nil {
		global.App.Log.Fatal(cons.FATAL_DB_CONNECT, zap.Any("err", err))
	}

	return db
}

// 初始化 mysql gorm.DB
func initMysqlGorm() (*gorm.DB, error) {
	return initGorm(cons.DATABASE_TYPE_MYSQL)
}

func initPostgresGorm() (*gorm.DB, error) {
	return initGorm(cons.DATABASE_TYPE_POSTGRESQL)
}

func initGorm(dbType string) (*gorm.DB, error) {

	// 根据数据库类型选择配置
	var dsn string

	// 从全局配置中获取数据库配置
	switch dbType {
	case cons.DATABASE_TYPE_MYSQL:
		// 获取 MySQL 数据库配置
		dbConfig := global.App.Config.MysqlDB

		// 是否有数据库配置
		if dbConfig.Database == "" {
			// 没有配置则直接返回,数据库配置错误
			return nil, fmt.Errorf(cons.ERROR_DB_CONFIG_DBNAME)
		}

		// 数据库配置
		dsn = fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
			dbConfig.UserName, cons.COLON, dbConfig.Password, cons.DATABASE_AT_TCP, cons.LEFT_ROUND_BRACKET,
			dbConfig.Host, cons.COLON, strconv.Itoa(dbConfig.Port), cons.RIGHT_ROUND_BRACKET, cons.SLASH,
			dbConfig.Database, cons.QUESTION_MARK, cons.DATABASE_CHARSET, cons.EQUAL, dbConfig.Charset, cons.DATABASE_MYSQL_PARAMS)

		mysqlConfig := mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         191,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 dateTime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据版本自动配置
		}

		db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,                                    // 禁用自动创建外键约束
			Logger:                                   getGormLogger(cons.DATABASE_TYPE_MYSQL), // 使用自定义 Logger
		})
		if err != nil {
			global.App.Log.Error(cons.ERROR_MYSQL_DB_CONNECT, zap.Any("err", err))
			return nil, err
		}

		// 连接池配置
		configConnectionPool(db, cons.DATABASE_TYPE_MYSQL)

		return db, nil

	case cons.DATABASE_TYPE_POSTGRESQL:
		// 获取 PostgreSQL 数据库配置
		dbConfig := global.App.Config.PostgresDB

		// 是否有数据库配置
		if dbConfig.Host == "" {
			// 没有配置则直接返回
			return nil, fmt.Errorf(cons.ERROR_DB_CONFIG_DBNAME)
		}

		// 数据库配置
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Dbname, dbConfig.Port, dbConfig.Sslmode, dbConfig.TimeZone)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,                                         // 禁用自动创建外键约束
			Logger:                                   getGormLogger(cons.DATABASE_TYPE_POSTGRESQL), // 使用自定义 Logger
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   cons.DATABASE_POSTGRESQL_TABLE_PREFIX, // 表名前缀，`Article` 的表名应该是 `t_articles`
				SingularTable: true,                                  // 使用单数表名，启用该选项后，`Article` 的表名应该是 `t_article`
			},
		})
		if err != nil {
			global.App.Log.Error(cons.ERROR_POSTGRES_DB_CONNECT, zap.Any("err", err))
			return nil, err
		}

		// 连接池配置
		configConnectionPool(db, cons.DATABASE_TYPE_POSTGRESQL)

		return db, nil

	default:
		return nil, fmt.Errorf(cons.ERROR_DB_TYPE_UNSUPPORT+cons.STRING_PLACEHOLDER, dbType)
	}
}

// 数据库连接池配置
func configConnectionPool(db *gorm.DB, dbType string) {

	// 根据数据库类型选择配置
	switch dbType {
	case cons.DATABASE_TYPE_MYSQL:
		// 获取 MySQL 数据库配置
		dbConfig := global.App.Config.MysqlDB

		// 连接池配置
		sqlDB, err := db.DB()
		if err != nil {
			global.App.Log.Error(cons.ERROR_MYSQL_DB_CONNECT, zap.Any("err", err))
			return
		}
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns) // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns) // SetMaxOpenConns 设置打开数据库连接的最大数量

	case cons.DATABASE_TYPE_POSTGRESQL:
		// 获取 PostgreSQL 数据库配置
		dbConfig := global.App.Config.PostgresDB

		// 连接池配置
		sqlDB, err := db.DB()
		if err != nil {
			global.App.Log.Error(cons.ERROR_POSTGRES_DB_CONNECT, zap.Any("err", err))
			return
		}
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns) // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns) // SetMaxOpenConns 设置打开数据库连接的最大数量

	default:
		fmt.Printf(cons.ERROR_DB_TYPE_UNSUPPORT+cons.STRING_PLACEHOLDER, dbType)
		return // 不支持的数据库类型
	}

	// 数据库表初始化
	initTable(db)
}

// 数据库表初始化
func initTable(db *gorm.DB) {

	// 自动迁移
	err := db.AutoMigrate(
		models.User{},
		models.Media{},
	)
	if err != nil {
		global.App.Log.Error(cons.ERROR_DB_MIGRATE, zap.Any("err", err))
		os.Exit(0)
	}
}

func getGormLogWriter(dbType string) logger.Writer {

	// 根据数据库类型选择日志文件名
	var writer io.Writer

	// 根据数据库类型选择日志文件名
	var logFilename string
	switch dbType {
	case cons.DATABASE_TYPE_MYSQL:
		if global.App.Config.MysqlDB.EnableFileLogWriter {
			logFilename = global.App.Config.MysqlDB.LogFilename
		}
	case cons.DATABASE_TYPE_POSTGRESQL:
		if global.App.Config.PostgresDB.EnableFileLogWriter {
			logFilename = global.App.Config.PostgresDB.LogFilename
		}
	}

	// 如果启用日志文件，则使用 lumberjack.Logger
	if logFilename != "" {
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + cons.SLASH + logFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}

func getGormLogger(dbType string) logger.Interface {

	// 根据配置文件设置日志等级
	var logMode logger.LogLevel

	switch dbType {
	case cons.DATABASE_TYPE_MYSQL:
		switch global.App.Config.MysqlDB.LogMode {
		case cons.LogLevelSilent:
			logMode = logger.Silent
		case cons.LogLevelError:
			logMode = logger.Error
		case cons.LogLevelWarn:
			logMode = logger.Warn
		case cons.LogLevelInfo:
			logMode = logger.Info
		default:
			logMode = logger.Info
		}
	case cons.DATABASE_TYPE_POSTGRESQL:
		switch global.App.Config.PostgresDB.LogMode {
		case cons.LogLevelSilent:
			logMode = logger.Silent
		case cons.LogLevelError:
			logMode = logger.Error
		case cons.LogLevelWarn:
			logMode = logger.Warn
		case cons.LogLevelInfo:
			logMode = logger.Info
		default:
			logMode = logger.Info
		}
	}

	return logger.New(
		getGormLogWriter(dbType),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logMode,                // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,                   // 彩色打印
		},
	)
}
