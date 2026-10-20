package usecase

import (
	"context"
	"log"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/app/model"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersSignIn(
	paramCtx context.Context,
	paramRequest *request.UsersAuthSignIn,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	//thisSpanCtx, thisSpan := receiverUseCase.trace.Span(paramCtx, "AuthSignIn")
	//thisSpan.SetAttributes(map[string]interface{}{
	//	"usecase.name":    "AuthSignIn",
	//	"usecase.request": paramRequest,
	//})
	//defer thisSpan.End()

	thisLang := headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))

	var (
		thisMessage       []messagepkg.Response
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

	// AdminSMSModel
	var requestForGet = request.UsersGet{Target: "mobile"}
	thisFetch, thisCount, thisErr := receiverUseCase.model.UsersGet(paramCtx, &requestForGet, thisMobileDetails.PhoneE164)
	if thisErr != nil {
		thisErrorResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"auth_sign_in",
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

	if thisFetch.Status == uint(model.Pending) {
		thisMessage = append(thisMessage, messagepkg.Pending(&thisLang, messagepkg.Keyword.Field["account"]))

		return responsepkg.Code403(
			thisMessage,
			nil, // paramRequest
			responsepkg.EntityV1UseCaseFeature{
				Api: responsepkg.EntityV1UseCaseApi{
					WithMessage: true,
				},
			},
		)
	}

	if thisFetch.Status == uint(model.Suspended) {
		thisMessage = append(thisMessage, messagepkg.Suspended(&thisLang, messagepkg.Keyword.Field["account"]))

		return responsepkg.Code403(
			thisMessage,
			nil, // paramRequest
			responsepkg.EntityV1UseCaseFeature{
				Api: responsepkg.EntityV1UseCaseApi{
					WithMessage: true,
				},
			},
		)
	}

	thisGenCode, _ := ghafpkg.GenCode( // Not error check
		receiverUseCase.config.OTP.Length,
		receiverUseCase.config.OTP.Chars,
	)
	thisLogicResponse, ok = receiverUseCase.redis.Set(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyOTP,
		thisMobileDetails.PhoneE164,
		thisGenCode,
		receiverUseCase.config.ExpireDuration.In3m,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	log.Println("OTP code is: ", thisGenCode)
	// Call sms service & send "OTP - thisCode(thisCode)" to user mobile

	thisMessage = append(thisMessage, messagepkg.Send(&thisLang, messagepkg.Keyword.Field["code"]))

	return responsepkg.Code200(
		thisMessage,
		nil, // response.Auth_sign_in{}
		// SuccessFull
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
