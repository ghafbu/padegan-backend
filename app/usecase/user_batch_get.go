package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersBatchGet(
	paramCtx context.Context,
	paramRequest *request.UsersBatchGet,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	var thisLang = headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	if paramRequest.Ids == nil &&
		paramRequest.Mobiles == nil &&
		paramRequest.Emails == nil &&
		paramRequest.NationalCodes == nil &&
		paramRequest.Phones == nil &&
		paramRequest.Usernames == nil &&
		paramRequest.Statuses == nil {
		// return error
		thisMessage = append(thisMessage, messagepkg.NotStandardRequest(&thisLang, messagepkg.Keyword.Section["admin_member"]))
		return responsepkg.Code400(
			thisMessage,
			paramRequest,
			responsepkg.EntityV1UseCaseFeature{
				Api: responsepkg.EntityV1UseCaseApi{
					WithMessage: true,
				},
			},
		)
	}

	thisFetch, thisCount, thisErr := receiverUseCase.model.UsersBatchGet(
		paramCtx, paramRequest)

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
	thisMessage = append(thisMessage, messagepkg.BatchGet(
		&thisLang, messagepkg.Keyword.Section["admin_member"], thisCount),
	)
	return responsepkg.Code200(
		thisMessage,
		response.UsersBatchGet{Users: thisFetch},
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
