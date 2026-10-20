package logics

import (
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func InArray(
	paramRequest interface{},
	paramRequestTarget *string,
	paramTargets map[string]string,
	paramLang string,
) (
	responsepkg.EntityV1UseCase,
	bool,
) {
	// Init Response Message
	var thisMessage []messagepkg.Response

	// Check Targets Exist
	if _, ok := paramTargets[*paramRequestTarget]; ok {
		return responsepkg.EntityV1UseCase{}, true
	} else {
		thisMessage = append(thisMessage, messagepkg.Match(&paramLang, messagepkg.Keyword.Field["id"]))
		return responsepkg.Code400(
			thisMessage,
			paramRequest,
		), true
	}
}
