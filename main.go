package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/pkg/snowflake"
	"web_app/routes"
	"web_app/settings"
	"web_app/translate"
)

func main() {
	//初始化配置学习
	err := settings.Init()
	if err != nil {
		fmt.Println("初始化配置信息失败 err=", err)
		return
	}

	err = logger.Init(settings.Conf.LogConfig, settings.Conf.Mode)
	if err != nil {
		fmt.Println("初始化全局日志失败 err=", err)
		return
	}
	zap.L().Sync() //延迟注册，加载缓冲区
	zap.L().Info(" 日志初始化成功")

	err = mysql.InitDB(settings.Conf.MySqlConfig)
	if err != nil {
		fmt.Println("初始化数据库连接失败 err=", err)
		return
	}
	defer mysql.Close()

	err = redis.Init(settings.Conf.RedisConfig)
	if err != nil {
		fmt.Println("初始化redis连接失败 err=", err)
		return
	}
	defer redis.Close()

	//注册雪花算法
	err = snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId)
	if err != nil {
		fmt.Println("初始化雪花算法异常 err", err)
		return
	}

	//注册翻译器
	if err = translate.InitTrans("zh"); err != nil {
		fmt.Println("初始化 国际化失败 err", err)
		return
	}

	r := routes.Setup(settings.Conf)
	fmt.Println("注册路由完成")

	//启动服务 平滑关机

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}
	fmt.Printf("web服务启动 %v", server.Addr)
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
