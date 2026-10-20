package usecase

import (
	"context"
	"log"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersResetPassword(
	paramCtx context.Context,
	paramRequest *request.UsersResetPassword,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	var thisLang = headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	// Logic: mobile details
	thisMobileDetails, thisLogicResponse, ok := logics.PhoneDetails(
		paramRequest, &paramRequest.Mobile, thisLang, "mobile",
	)
	if !ok {
		return thisLogicResponse
	}

	// Get user by mobile
	var requestForGet = request.UsersGet{Target: "mobile"}

	_, thisCount, thisErr := receiverUseCase.model.UsersGet(
		paramCtx,
		&requestForGet,
		thisMobileDetails.PhoneE164,
	)
	if thisErr != nil {
		thisErrorResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"reset_password",
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

	// thisCode, thisErr := receiverUseCase.otp.Set(thisMobileDetails.PhoneE164)
	// if thisErr != nil {
	//	thisErrorResponse := errors.OperationAction(
	//		thisLang,
	//		paramRequest,
	//		thisErr,
	//		"admin_members",
	//		"reset_password_OTP",
	//	)
	//	return thisErrorResponse
	// }

	thisGenCode, _ := ghafpkg.GenCode( // Not error check
		receiverUseCase.config.OTP.Length,
		receiverUseCase.config.OTP.Chars,
	)
	thisLogicResponse, ok = receiverUseCase.redis.Set(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyAdminJWTResetPassword,
		thisMobileDetails.PhoneE164,
		thisGenCode,
		receiverUseCase.config.ExpireDuration.In2m,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	log.Println("OTP code is: ", thisGenCode)
	// Call sms service & send "OTP - thisCode(thisCode)" to user mobile

	// SuccessFull
	thisMessage = append(thisMessage, messagepkg.Send(&thisLang, messagepkg.Keyword.Field["code"]))
	return responsepkg.Code200(
		thisMessage,
		nil, // response.UsersResetPassword{}
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
