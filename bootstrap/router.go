package bootstrap

import (
	"context"
	"go-gin/cons"
	"go-gin/global"
	"go-gin/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 前端项目静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	router.Static("/public", "./static")
	router.Static("/storage", "./storage/app/public")

	// 获取API前缀和版本
	apiPrefix := global.App.Config.Api.Prefix
	apiVersion := global.App.Config.Api.Version

	// 注册api分组路由
	apiGroup := router.Group(apiPrefix + cons.SLASH + apiVersion)
	routes.SetApiGroupRoutes(apiGroup)

	// 注册用户分组路由
	userGroup := router.Group(apiPrefix + cons.SLASH + apiVersion + cons.SLASH + cons.API_USER_GROUP)
	routes.SetUserGroupRoutes(userGroup)

	return router
}

// RunServer 优雅重启/停止服务器
func RunServer() {
	r := setupRouter()

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    cons.COLON + strconv.Itoa(global.App.Config.App.Port),
		Handler: r,
	}

	// 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(cons.FATAL_SERVER_START+cons.STRING_PLACEHOLDER_N, err)
		}
		log.Println(cons.INFO_SERVER_START + strconv.Itoa(global.App.Config.App.Port))
	}()

	/**
	 *等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	 */
	quit := make(chan os.Signal, 1) // 创建一个带缓冲的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println(cons.INFO_SERVER_IN_SHUTDOWN)

	// 设置 5 秒的超时时间来优雅地关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(cons.FATAL_SERVER_SHUTDOWN, err)
	}
	log.Println(cons.INFO_SERVER_SHUTDOWN)
}
