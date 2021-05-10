// description: 日志
//
// author: vignetting
// time: 2021/5/10

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"structure/pkg/setting"
	"time"
)

var sugaredLogger *zap.SugaredLogger
var logger *zap.Logger

func Setup() {
	encoderConfig := zapcore.EncoderConfig{
		// key 名称
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",

		// 各类型数据的编码器
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	// 日志分割配置
	hook := lumberjack.Logger{
		// 日志存放路径
		Filename: setting.LogSetting.FilePath,

		// 每个日志文件保存的最大尺寸，单位：M
		MaxSize: setting.LogSetting.FileMaxSize,

		// 旧文件最多被保留的数目，0 表示都保留
		MaxBackups: setting.LogSetting.MaxBackups,

		// 文件最多保存多长时间，单位是天
		MaxAge: setting.LogSetting.MaxRetentionTime,

		// 是否压缩
		Compress: setting.LogSetting.Compress,
	}

	// 构建日志核心
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		getLevel(setting.LogSetting.Level),
	)

	// 构造日志，会显示调用者
	if setting.LogSetting.DeveloperMode {
		logger = zap.New(core, zap.AddCaller(), zap.Development())
		sugaredLogger = logger.Sugar()
	} else {
		logger = zap.New(core, zap.AddCaller())
		sugaredLogger = logger.Sugar()
	}

	// 修改 zap 中的默认日志
	zap.ReplaceGlobals(logger)
}

func Debug(logMessage ...interface{}) {
	sugaredLogger.Debug(logMessage)
}

func Info(logMessage ...interface{}) {
	sugaredLogger.Info(logMessage)
}

func Warn(logMessage ...interface{}) {
	sugaredLogger.Warn(logMessage)
}

func Error(logMessage ...interface{}) {
	sugaredLogger.Error(logMessage)
}

func DPanic(logMessage ...interface{}) {
	sugaredLogger.DPanic(logMessage)
}

func Panic(logMessage ...interface{}) {
	sugaredLogger.Panic(logMessage)
}

func fatal(logMessage ...interface{}) {
	sugaredLogger.Fatal(logMessage)
}

func Logger() *zap.Logger {
	return logger
}

func getLevel(level string) zap.AtomicLevel {
	atomicLevel := zap.NewAtomicLevel()
	switch level {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "error":
		atomicLevel.SetLevel(zap.WarnLevel)
	case "dpanic":
		atomicLevel.SetLevel(zap.DPanicLevel)
	case "panic":
		atomicLevel.SetLevel(zap.PanicLevel)
	case "fatal":
		atomicLevel.SetLevel(zap.FatalLevel)
	default:
		panic("错误的日志级别")
	}

	return atomicLevel
}
