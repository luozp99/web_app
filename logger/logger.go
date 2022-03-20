package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
	"web_app/settings"
)

func Init(conf *settings.LogConfig, mode string) (err error) {
	writerSyncer := getLogWriter(conf)
	encoder := getEncoder()

	var level = new(zapcore.Level)
	err = level.UnmarshalText([]byte(conf.Level))
	if err != nil {
		return err
	}

	var core zapcore.Core

	if mode == "dev" {
		cosoleEnbcoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writerSyncer, level),
			zapcore.NewCore(cosoleEnbcoder, zapcore.Lock(os.Stdout), level),
		)

	} else {
		core = zapcore.NewCore(encoder, writerSyncer, level)
	}
	logger := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logger.Sugar().Desugar())
	return err
}

func getEncoder() zapcore.Encoder {

	encoder := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewJSONEncoder(encoder)
}

func getLogWriter(conf *settings.LogConfig) zapcore.WriteSyncer {
	//file, _ := os.OpenFile("./test.log",os.O_CREATE|os.O_APPEND|os.O_RDONLY,0744)
	lumberjackLogger := &lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackUps,
		MaxAge:     conf.MaxAge,
		Compress:   false,
	}

	return zapcore.AddSync(lumberjackLogger)
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(start)

		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path, zap.Any("error", err), zap.String("request", string(httpRequest)))
					c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())))
				} else {
					zap.L().Error("Recovery from panic", zap.Any("error", err), zap.String("request", string(httpRequest)))
				}
			}
		}()
		c.Next()
	}
}
