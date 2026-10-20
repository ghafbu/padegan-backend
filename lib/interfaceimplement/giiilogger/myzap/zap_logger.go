package myzap

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/dromara/carbon/v2"
	"github.com/ghafbu/padegan-backend/config"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/google/uuid"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapSinLogger *zap.SugaredLogger
	once         = sync.Once{}
)

type MyZapLogger struct {
	cfg    config.Logger
	logger *zap.SugaredLogger
}

func Init(cfg config.Logger) *MyZapLogger {
	logger := &MyZapLogger{cfg: cfg}
	logger.setup()

	return logger
}

func (l *MyZapLogger) Debug(paramTags []string, paramData map[string]interface{}, paramDescription string) {
	params := generateLogMap(paramTags, paramData)
	l.logger.Debugw(paramDescription, params...)
}

func (l *MyZapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *MyZapLogger) Info(paramTags []string, paramData map[string]interface{}, paramDescription string) {
	params := generateLogMap(paramTags, paramData)
	l.logger.Infow(paramDescription, params...)
}

func (l *MyZapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *MyZapLogger) Panic(paramTags []string, paramData map[string]interface{}, paramDescription string) {
	params := generateLogMap(paramTags, paramData)
	l.logger.Panicw(paramDescription, params...)
}

func (l *MyZapLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args)
}

func (l *MyZapLogger) Warn(paramTags []string, paramData map[string]interface{}, paramDescription string) {
	params := generateLogMap(paramTags, paramData)
	l.logger.Warnw(paramDescription, params...)
}

func (l *MyZapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *MyZapLogger) Error(paramTags []string, paramData map[string]interface{}, paramDescription string) {
	params := generateLogMap(paramTags, paramData)
	l.logger.Errorw(paramDescription, params...)
}

func (l *MyZapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *MyZapLogger) Fatal(paramTags []string, paramData map[string]interface{}, paramDescription string) {
	params := generateLogMap(paramTags, paramData)
	l.logger.Fatalw(paramDescription, params...)
}

func (l *MyZapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

var zapLogLevelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func (l *MyZapLogger) setup() {
	once.Do(func() {
		englishDate := time.Now().Format("2006-01-02 15:04:05")
		jalaliDate := carbon.Parse(englishDate).Persian().String()
		jalaliDateOnly := jalaliDate[:10]

		// بررسی اینکه مسیر لاگ یک فایل نباشد
		if fileInfo, err := os.Stat(l.cfg.FilePath); err == nil && !fileInfo.IsDir() {
			log.Fatalf("zap-logger: log path exists but is a file, not a directory: %s", l.cfg.FilePath)
		}

		// بررسی و ایجاد مسیر پوشه لاگ‌ها
		if err := os.MkdirAll(l.cfg.FilePath, 0755); err != nil {
			log.Fatalf("zap-logger: failed to create log directory: %v", err)
		}

		var thisFileName = fmt.Sprintf("%s-%s.%s", jalaliDateOnly, uuid.New(), "log")
		thisFileUrl := filepath.Join(l.cfg.FilePath, thisFileName)

		// fmt.Println("zap-logger:  setup thisFileUrl:", thisFileUrl)
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   thisFileUrl,
			MaxSize:    1,
			MaxAge:     20,
			LocalTime:  true,
			MaxBackups: 5,
			Compress:   true,
		})

		config := zap.NewProductionEncoderConfig()
		// config.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncodeTime = jalaliTimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			l.getLogLevel(),
		)

		logger := zap.New(core, zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zapcore.ErrorLevel),
		).Sugar()

		// zapSinLogger = logger.With("ServiceName", "user", "LoggerName", "Zaplog")
		zapSinLogger = logger.With("ts_solar", jalaliDate, "LoggerProvider", "MyZapLogger")
	})

	l.logger = zapSinLogger
}

func prepareLogInfo(cat string, sub string, child string, extra map[string]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[string]interface{})
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	extra["ChildCategory"] = child

	return logParamsToZapParams(extra)
}

func generateLogMap(paramTags []string, paramData map[string]interface{}) []interface{} {
	paramNewData := make(map[string]interface{})
	paramNewData["tags"] = paramTags
	paramNewData["data"] = paramData

	return logParamsToZapParams(paramNewData)
}

func (l *MyZapLogger) getLogLevel() zapcore.Level {
	level, exists := zapLogLevelMapping[l.cfg.Logger]
	if !exists {
		return zapcore.DebugLevel
	}

	return level
}
