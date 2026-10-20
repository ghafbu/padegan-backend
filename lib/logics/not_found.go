package logics

import (
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func NotFound(
	paramRequest interface{},
	paramCount int64,
	paramLang string,
	paramSectionName string,
) (
	responsepkg.EntityV1UseCase,
	bool,
) {
	var thisMessage []messagepkg.Response

	if paramCount == 0 {
		thisMessage = append(thisMessage, messagepkg.NotFoundSection(&paramLang, messagepkg.Keyword.Section[paramSectionName]))
		return responsepkg.Code404(
			thisMessage,
			paramRequest,
		), false
	}

	return responsepkg.EntityV1UseCase{}, true
}
