package usecase

import (
	"context"
	"github.com/ghafbu/padegan-backend/lib/errors"

	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UserSettingsUseCase) UserSettingsGet(
	paramCtx context.Context,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	var thisLang = headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))

	var thisMessage []messagepkg.Response
	// Modal
	thisFetch, thisCount, thisErr := receiverUseCase.model.UserSettingsGet(paramCtx)
	if thisErr != nil {
		thisLogicResponse := errors.ModelAction(
			thisErr,
			nil,
			thisLang,
			"admin_settings_get",
			"get",
		)

		return thisLogicResponse
	}

	// Logic: Not found
	thisLogicResponse, thisFound := logics.NotFound(
		nil, thisCount, thisLang, "admin_settings_get",
	)
	if !thisFound {
		return thisLogicResponse
	}

	// SuccessFull
	thisMessage = append(thisMessage, messagepkg.Get(&thisLang, messagepkg.Keyword.Section["admin_settings"]))
	return responsepkg.Code200(
		thisMessage,
		thisFetch,
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
