package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersBatchUpdate(
	paramCtx context.Context,
	paramRequest *request.UsersBatchUpdate,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	var thisLang = headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	thisFetch, thisCount, thisErr := receiverUseCase.model.UsersBatchUpdate(
		paramCtx, paramRequest,
	)

	if thisErr != nil {
		thisLogicResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"batch_get",
		)
		return thisLogicResponse
	}

	// Check access

	// Logic: Not found
	thisLogicResponse, thisFound := logics.NotFound(
		paramRequest, thisCount, thisLang, "admin_member",
	)
	if !thisFound {
		return thisLogicResponse
	}

	// SuccessFull
	thisMessage = append(thisMessage, messagepkg.BatchUpdate(
		&thisLang, messagepkg.Keyword.Section["admin_member"], thisCount),
	)
	return responsepkg.Code200(
		thisMessage,
		thisFetch, // thisFetch or below code
		// response.UsersBatchUpdate{
		//	Users: thisFetch.Users,
		//	Paper:        thisFetch.Paper,
		// }
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
