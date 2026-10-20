package myzap

import (
	"time"

	"go.uber.org/zap/zapcore"
)

func jalaliTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	englishDate := t.Format("2006-01-02 15:04:05")
	enc.AppendString(englishDate)
}
