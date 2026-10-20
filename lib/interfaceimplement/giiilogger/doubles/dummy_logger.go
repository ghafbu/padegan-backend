package doubles

type DummyLogger struct {
}

func NewDummyLogger() *DummyLogger {
	return &DummyLogger{}
}

func (d DummyLogger) Debug(_ []string, _ map[string]interface{}, _ string) {
}

func (d DummyLogger) Debugf(_ string, _ ...interface{}) {
}

func (d DummyLogger) Info(_ []string, _ map[string]interface{}, _ string) {
}

func (d DummyLogger) Infof(_ string, _ ...interface{}) {
}

func (d DummyLogger) Panic(_ []string, _ map[string]interface{}, _ string) {
}

func (d DummyLogger) Panicf(_ string, _ ...interface{}) {
}

func (d DummyLogger) Warn(_ []string, _ map[string]interface{}, _ string) {
}

func (d DummyLogger) Warnf(_ string, _ ...interface{}) {
}

func (d DummyLogger) Error(_ []string, _ map[string]interface{}, _ string) {
}

func (d DummyLogger) Errorf(_ string, _ ...interface{}) {
}

func (d DummyLogger) Fatalf(_ string, _ ...interface{}) {
}
