package log

import (
	"applet-server/internal/conf"
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var defaultLogger *MyLogger
var ProviderSet = wire.NewSet(NewLogger)

type MyLogger struct {
	logger     *zap.Logger
	logSetting *conf.Log
	hook       *lumberjack.Logger
	// 日期
	date       string
	messageKey string
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

func NewLogger(confLog *conf.Log) *MyLogger {
	myLogger := &MyLogger{
		logSetting: confLog,
		messageKey: "msg",
	}
	logLevel := strings.ToLower(confLog.Level)

	var syncer zapcore.WriteSyncer
	if confLog.LogInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(myLogger))
	} else {
		syncer = zapcore.AddSync(myLogger)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "log",
		CallerKey:     "line",
		StacktraceKey: "stacktrace",

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

	myLogger.logger = logger.WithOptions(zap.AddCallerSkip(2))
	defaultLogger = myLogger
	return myLogger
}

func (l *MyLogger) WithContext(ctx context.Context) {
	traceId := ctx.Value("trace_id").(string)
	if traceId != "" {
		l.logger.With(zap.String("trace_id", ctx.Value("trace_id").(string)))
	}
	sessionId := ctx.Value("session_id").(string)
	if sessionId != "" {
		l.logger.With(zap.String("session_id", ctx.Value("session_id").(string)))
	}
}

func (l *MyLogger) Log(level log.Level, keyvals ...interface{}) error {
	keyLen := len(keyvals)
	if keyLen == 0 || keyLen%2 != 0 {
		l.logger.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	data := make([]zap.Field, 0, (keyLen/2)+1)
	for i := 0; i < keyLen; i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.logger.Debug("", data...)
	case log.LevelInfo:
		l.logger.Info("", data...)
	case log.LevelWarn:
		l.logger.Warn("", data...)
	case log.LevelError:
		l.logger.Error("", data...)
	case log.LevelFatal:
		l.logger.Fatal("", data...)
	}
	return nil
}

// Debug logs a message at debug level.
func (l *MyLogger) Debug(a ...interface{}) {
	l.Log(log.LevelDebug, l.messageKey, a)
}

func Debug(a ...interface{}) {
	defaultLogger.Log(log.LevelDebug, defaultLogger.messageKey, a)
}

// Debugf logs a message at debug level.
func (l *MyLogger) Debugf(format string, a ...interface{}) {
	_ = l.Log(log.LevelDebug, l.messageKey, fmt.Sprintf(format, a...))
}

func Debugf(format string, a ...interface{}) {
	defaultLogger.Log(log.LevelDebug, defaultLogger.messageKey, fmt.Sprintf(format, a...))
}

// Debug logs a message at debug level.
func (l *MyLogger) Info(a ...interface{}) {
	l.Log(log.LevelInfo, l.messageKey, a)
}

func Info(a ...interface{}) {
	defaultLogger.Log(log.LevelInfo, defaultLogger.messageKey, a)
}

// Debugf logs a message at debug level.
func (l *MyLogger) Infof(format string, a ...interface{}) {
	_ = l.Log(log.LevelInfo, l.messageKey, fmt.Sprintf(format, a...))
}

func Infof(format string, a ...interface{}) {
	defaultLogger.Log(log.LevelInfo, defaultLogger.messageKey, fmt.Sprintf(format, a...))
}

// Debug logs a message at debug level.
func (l *MyLogger) Warn(a ...interface{}) {
	l.Log(log.LevelWarn, l.messageKey, a)
}

func Warn(a ...interface{}) {
	defaultLogger.Log(log.LevelWarn, defaultLogger.messageKey, a)
}

// Debugf logs a message at debug level.
func (l *MyLogger) Warnf(format string, a ...interface{}) {
	_ = l.Log(log.LevelWarn, l.messageKey, fmt.Sprintf(format, a...))
}

// Debug logs a message at debug level.
func (l *MyLogger) Error(a ...interface{}) {
	l.Log(log.LevelError, l.messageKey, a)
}

func Error(a ...interface{}) {
	defaultLogger.Log(log.LevelError, defaultLogger.messageKey, a)
}

// Debugf logs a message at debug level.
func (l *MyLogger) Errorf(format string, a ...interface{}) {
	_ = l.Log(log.LevelError, l.messageKey, fmt.Sprintf(format, a...))
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
