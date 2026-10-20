package giiilogger

import (
	"github.com/ghafbu/padegan-backend/config"
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiilogger/myzap"
)

const (
	ImplementNameZap = "myZap"
)

func NewILogger(paramImplementName string, config config.Logger) _interface.ILogger {
	switch paramImplementName {
	case ImplementNameZap:
		return myzap.Init(config)
	default:
		return myzap.Init(config)
	}
}
