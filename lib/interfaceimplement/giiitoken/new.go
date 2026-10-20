package giiitoken

import (
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiitoken/myjwt"
)

const (
	ImplementNameJWT = "myJWT"
)

func NewIToken(paramImplementName string) _interface.IToken {
	switch paramImplementName {
	case ImplementNameJWT:
		return myjwt.Init()
	default:
		return myjwt.Init()
	}
}
