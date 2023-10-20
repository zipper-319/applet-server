package MyLog

import (
	"applet-server/internal/conf"
	"errors"
	"fmt"
	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var ProviderSet = wire.NewSet(NewLogger)

type MyLogger struct {
	logger     *kratoszap.Logger
	logSetting *conf.Log
	hook       *lumberjack.Logger
	// 日期
	date string
}

func (l *MyLogger) Write(data []byte) (n int, err error) {
	if l == nil || l.logSetting == nil {
		return 0, errors.New("MyLogger is nil")
	}

	dateTime := time.Now().Format(l.logSetting.TimeFormat)
	if l.hook == nil {
		filePath := l.getLogFilePath()
		fileName := l.getLogFileName(dateTime)
		l.hook = &lumberjack.Logger{
			Filename:   filePath + "/" + fileName,    // 日志文件路径
			MaxSize:    int(l.logSetting.MaxSize),    // megabytes
			MaxBackups: int(l.logSetting.MaxBackups), // 最多保留300个备份
			Compress:   l.logSetting.Compress,        // 是否压缩 disabled by default
		}
		maxAge := int(l.logSetting.MaxDays)
		if maxAge > 0 {
			l.hook.MaxAge = maxAge // days
		}
		l.date = dateTime
	}
	n, e := l.hook.Write(data)
	//按照每天生成日志文件
	if l.date != dateTime {
		filePath := l.getLogFilePath()
		fileName := l.getLogFileName(dateTime)
		l.hook.Filename = filePath + "/" + fileName
	}

	return n, e
}

func NewLogger(confLog *conf.Log) log.Logger {
	myLogger := &MyLogger{
		logSetting: confLog,
	}
	logLevel := strings.ToLower(confLog.Level)

	var syncer zapcore.WriteSyncer
	if confLog.LogInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(myLogger))
	} else {
		syncer = zapcore.AddSync(myLogger)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "line",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	if confLog.JsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		level,
	)

	logger := zap.New(core)

	if confLog.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	logger = logger.WithOptions(zap.AddCallerSkip(2))
	myLogger.logger = kratoszap.NewLogger(logger)
	defaultLogger := myLogger.logger
	log.SetLogger(defaultLogger)
	return myLogger.logger
}

// getLogFilePath get the log file save path
func (logger *MyLogger) getLogFilePath() string {
	return fmt.Sprintf("%s%s", logger.logSetting.GetRootPath(), logger.logSetting.GetSavePath())
}

// getLogFileName get the save name of the log file
func (logger *MyLogger) getLogFileName(dateTime string) string {
	return fmt.Sprintf("%s%s.log",
		logger.logSetting.GetSaveFilename(),
		dateTime,
	)
}

type GormLogger struct {
	name   string
	logger *zap.Logger
}

func NewGormLogger() *GormLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &GormLogger{
		name:   "gorm_logger",
		logger: logger,
	}
}

func (g *GormLogger) Print(values ...interface{}) {
	if len(values) < 2 {
		return
	}

	switch values[0] {
	case "sql":
		g.logger.Debug("gorm.debug.sql",
			zap.String("query", values[3].(string)),
			zap.Any("values", values[4]),
			zap.Float64("duration in ms", float64(values[2].(time.Duration))/float64(time.Millisecond)),
			zap.Int64("affected-rows", values[5].(int64)),
			zap.String("source", values[1].(string)), // if AddCallerSkip(6) is well defined, we can safely remove this field
		)
	default:
		g.logger.Debug("gorm.debug.other",
			zap.Any("values", values[2:]),
			zap.String("source", values[1].(string)), // if AddCallerSkip(6) is well defined, we can safely remove this field
		)
	}

}
