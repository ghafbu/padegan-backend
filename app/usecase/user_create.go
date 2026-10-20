package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/lib/logics"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersCreate(
	paramCtx context.Context,
	paramRequest *request.UsersCreate,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	thisLang := headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var (
		thisMessage       []messagepkg.Response
		thisPhoneDetails  phonepkg.PhoneDetails
		thisMobileDetails phonepkg.PhoneDetails
		thisLogicResponse responsepkg.EntityV1UseCase
		ok                bool
	)

	// Logic: mobile details
	thisMobileDetails, thisLogicResponse, ok = logics.PhoneDetails(
		paramRequest, &paramRequest.Mobile, thisLang, "mobile",
	)
	if !ok {
		return thisLogicResponse
	}

	// Logic: phone details
	if paramRequest.Phone != nil { // Check, because pointer nil
		thisPhoneDetails, thisLogicResponse, ok = logics.PhoneDetails(
			paramRequest, paramRequest.Phone, thisLang, "phone",
		)
		if !ok {
			return thisLogicResponse
		}
	}

	// Model
	thisFetch, thisErr := receiverUseCase.model.UsersCreate(
		paramCtx, paramRequest, thisMobileDetails, thisPhoneDetails,
	)
	if thisErr != nil {
		thisErrorResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"create",
		)
		return thisErrorResponse
	}

	thisMessage = append(thisMessage, messagepkg.Create(&thisLang, messagepkg.Keyword.Section["admin_member"]))
	return responsepkg.Code201(
		thisMessage,
		thisFetch, // thisFetch = response.public
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
