package usecase

import (
	"context"
	"github.com/ghafbu/padegan-backend/lib/errors"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UserSettingsUseCase) UserSettingsUpdate(
	paramCtx context.Context,
	paramRequest *request.UserSettingsUpdate,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	thisLang := headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	// Model
	thisFetch, thisErr := receiverUseCase.model.UserSettingsUpdate(paramCtx, paramRequest)
	if thisErr != nil {
		thisLogicResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_settings_update",
			"update",
		)
		return thisLogicResponse
	}

	// Success update
	thisMessage = append(thisMessage, messagepkg.Update(&thisLang, messagepkg.Keyword.Section["admin_settings"]))
	return responsepkg.Code200(
		thisMessage,
		response.UserSettingsUpdate{Settings: thisFetch},
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
