/**
 * 直接复制的 github.com/dubbogo/gost/log/logger，做了很多修改
 */

package logger

import (
	"github.com/natefinch/lumberjack"

	gostLogger "github.com/dubbogo/gost/log/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var dubboLogger *DubboLogger

func init() {
	InitLogger(nil)
}

// nolint
type DubboLogger struct {
	gostLogger.Logger
	dynamicLevel zap.AtomicLevel
	ZapLogger    *zap.Logger
}

type Config struct {
	LumberjackConfig *lumberjack.Logger `yaml:"lumberjack-config"`
	ZapConfig        *zap.Config        `yaml:"zap-config"`
	CallerSkip       int
}

// InitLogger use for init logger by @conf
func InitLogger(conf *Config) {
	var (
		zapLogger *zap.Logger
		config    = &Config{}
	)
	if conf == nil || conf.ZapConfig == nil {
		zapLoggerEncoderConfig := GetZapEncoderConfigDefault()
		config.ZapConfig = GetZapConfigDefault(zapLoggerEncoderConfig)
	} else {
		config.ZapConfig = conf.ZapConfig
	}

	if conf != nil {
		config.CallerSkip = conf.CallerSkip
	}

	if config.CallerSkip == 0 {
		config.CallerSkip = 1
	}

	if conf == nil || conf.LumberjackConfig == nil {
		zapLogger, _ = config.ZapConfig.Build(zap.AddCaller(), zap.AddCallerSkip(config.CallerSkip))
	} else {
		config.LumberjackConfig = conf.LumberjackConfig
		zapLogger = initZapLoggerWithSyncer(config)
	}

	dubboLogger = &DubboLogger{Logger: zapLogger.Sugar(), dynamicLevel: config.ZapConfig.Level, ZapLogger: zapLogger}

	SetGlobalLogger(dubboLogger)
}

func GetZapConfigDefault(encoderConfig zapcore.EncoderConfig) *zap.Config {
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func GetZapEncoderConfigDefault() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// SetGlobalLogger 设置框架全局 logger
func SetGlobalLogger(dubboLogger *DubboLogger) {
	gostLogger.SetLogger(dubboLogger)
}

func GetDubboLogger() *DubboLogger {
	return dubboLogger
}

func GetLogger() gostLogger.Logger {
	return dubboLogger
}

// SetLoggerLevel use for set logger level
func SetLoggerLevel(level string) bool {
	if l, ok := dubboLogger.Logger.(OpsLogger); ok {
		l.SetLoggerLevel(level)
		return true
	}
	return false
}

// OpsLogger use for the SetLoggerLevel
type OpsLogger interface {
	gostLogger.Logger
	SetLoggerLevel(level string)
}

// SetLoggerLevel use for set logger level
func (dl *DubboLogger) SetLoggerLevel(level string) {
	l := new(zapcore.Level)
	if err := l.Set(level); err == nil {
		dl.dynamicLevel.SetLevel(*l)
	}
}

// initZapLoggerWithSyncer init zap Logger with syncer
func initZapLoggerWithSyncer(conf *Config) *zap.Logger {
	core := zapcore.NewCore(
		conf.getEncoder(),
		conf.getLogWriter(),
		zap.NewAtomicLevelAt(conf.ZapConfig.Level.Level()),
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(conf.CallerSkip))
}

// getEncoder get encoder by config, zapcore support json and console encoder
func (c *Config) getEncoder() zapcore.Encoder {
	if c.ZapConfig.Encoding == "json" {
		return zapcore.NewJSONEncoder(c.ZapConfig.EncoderConfig)
	} else if c.ZapConfig.Encoding == "console" {
		return zapcore.NewConsoleEncoder(c.ZapConfig.EncoderConfig)
	}
	return nil
}

// getLogWriter get Lumberjack writer by LumberjackConfig
func (c *Config) getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(c.LumberjackConfig)
}

func (l *DubboLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *DubboLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *DubboLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l *DubboLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l *DubboLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *DubboLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *DubboLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *DubboLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *DubboLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l *DubboLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}
