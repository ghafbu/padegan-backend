package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/api/dto/params"
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersUpdate(
	paramCtx context.Context,
	paramRequest *request.UsersUpdate,
	paramHeader map[string][]string,
	paramParams params.UsersUpdate,
) responsepkg.EntityV1UseCase {
	thisLang := headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var (
		thisMessage       []messagepkg.Response
		thisPhoneDetails  phonepkg.PhoneDetails
		thisLogicResponse responsepkg.EntityV1UseCase
		ok                bool
	)

	// Logic: phone details
	if paramRequest.Phone != nil { // Check, because pointer nil
		thisPhoneDetails, thisLogicResponse, ok = logics.PhoneDetails(
			paramRequest, paramRequest.Phone, thisLang, "phone",
		)
		if !ok {
			return thisLogicResponse
		}
	}

	// Logic: mobile details
	// thisMobileDetails, thisLogicResponse, ok = logics.PhoneDetails(
	// 	paramRequest, &paramRequest.Mobile, thisLang, "mobile",
	// )
	// if !ok {
	// 	return thisLogicResponse
	// }

	thisFetch, thisCount, thisErr := receiverUseCase.model.UsersUpdate(
		paramCtx, paramRequest, thisPhoneDetails, paramParams,
	)
	if thisErr != nil {
		thisErrorResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"update",
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

	// Success update
	thisMessage = append(thisMessage, messagepkg.Update(&thisLang, messagepkg.Keyword.Section["admin_member"]))
	return responsepkg.Code200(
		thisMessage,
		response.UsersUpdate{Member: thisFetch},
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
