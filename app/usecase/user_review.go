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

func (receiverUseCase *UsersUseCase) UsersReview(
	paramCtx context.Context,
	paramRequest *request.UsersReview,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	var thisLang = headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	thisFetch, thisCount, thisErr := receiverUseCase.model.UsersReview(
		paramCtx, paramRequest)
	if thisErr != nil {
		thisErrorResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"review",
		)
		return thisErrorResponse
	}

	// Logic: Not found
	thisLogicResponse, thisFound := logics.NotFound(
		paramRequest, thisCount, thisLang, "admin_member",
	)
	if !thisFound {
		return thisLogicResponse
	}

	// Successful
	thisMessage = append(thisMessage, messagepkg.Search(&thisLang, messagepkg.Keyword.Section["admin_member"], thisCount))
	return responsepkg.Code200(
		thisMessage,
		thisFetch, // thisFetch or below code
		// response.UsersReview{
		//	Users: thisFetch.Users,
		//	Paper:        thisFetch.Paper,
		// },
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
