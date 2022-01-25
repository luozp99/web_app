package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/setting"
)

func main() {
	//初始化配置学习
	err := setting.Init()
	if err != nil {
		fmt.Println("初始化配置信息失败 err=", err)
		return
	}

	err = logger.Init()
	if err != nil {
		fmt.Println("初始化全局日志失败 err=", err)
		return
	}
	zap.L().Sync() //延迟注册，加载缓冲区
	zap.L().Info(" 日志初始化成功")

	err = mysql.InitDB()
	if err != nil {
		fmt.Println("初始化数据库连接失败 err=", err)
		return
	}
	defer mysql.Close()

	err = redis.Init()
	if err != nil {
		fmt.Println("初始化redis连接失败 err=", err)
		return
	}
	defer redis.Close()

	r := routes.Setup()
	fmt.Println("注册路由完成")

	//启动服务 平滑关机

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("监听服务关闭", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit // 阻塞此处，当有接收到信号之后才进行往下执行

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server shutdown :", zap.Error(err))
	}

	zap.L().Info("server exit~~~")

}
