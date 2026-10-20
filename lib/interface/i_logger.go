package _interface

type ILogger interface {
	Debug(paramTags []string, paramData map[string]interface{}, paramDescription string)
	Debugf(template string, args ...interface{})

	Info(paramTags []string, paramData map[string]interface{}, paramDescription string)
	Infof(template string, args ...interface{})

	Panic(paramTags []string, paramData map[string]interface{}, paramDescription string)
	Panicf(template string, args ...interface{})

	Warn(paramTags []string, paramData map[string]interface{}, paramDescription string)
	Warnf(template string, args ...interface{})

	Error(paramTags []string, paramData map[string]interface{}, paramDescription string)
	Errorf(template string, args ...interface{})

	Fatalf(template string, args ...interface{})
}
