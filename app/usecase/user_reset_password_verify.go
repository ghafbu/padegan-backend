package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersResetPasswordVerify(
	paramCtx context.Context,
	paramRequest *request.UsersResetPasswordVerify,
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

	// Verify
	// thisIsEqual, thisIsExist, thisErr := receiverUseCase.otp.Verify(thisMobileDetails.PhoneE164, paramRequest.Code)
	// if thisErr != nil {
	//	thisErrorResponse := errors.OperationAction(
	//		thisLang,
	//		paramRequest,
	//		thisErr,
	//		"admin_members",
	//		"reset_password_verify_OTP",
	//	)
	//	return thisErrorResponse
	// }
	//
	// // Is not exist
	// if !thisIsExist {
	//	thisMessage = append(thisMessage, messagepkg.NotFoundSection(&thisLang, messagepkg.Keyword.Section["admin_member"]))
	//	return responsepkg.Code404(
	//		thisMessage,
	//		paramRequest,
	//		responsepkg.EntityV1UseCaseFeature{
	//			Api: responsepkg.EntityV1UseCaseApi{},
	//		},
	//	)
	// }
	//
	// // Is not true(code not equal code in redis)
	// if !thisIsEqual {
	//	thisMessage = append(thisMessage, messagepkg.Wrong(&thisLang, messagepkg.Keyword.Field["code"]))
	//	return responsepkg.Code401(
	//		thisMessage,
	//		paramRequest,
	//		responsepkg.EntityV1UseCaseFeature{
	//			Api: responsepkg.EntityV1UseCaseApi{
	//				WithMessage: true,
	//			},
	//		},
	//	)
	// }

	// thisLogicResponse, thisCodeFromRedis, ok := logics.RedisGetWithDecrypt(
	thisLogicResponse, thisCodeFromRedis, ok := receiverUseCase.redis.Get(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyAdminJWTResetPassword,
		thisMobileDetails.PhoneE164,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	// thisLogicResponse, customClaims, ok := logics.TokenParse[string](
	// 	paramRequest,
	// 	thisLang,
	// 	receiverUseCase.cfg.Token.ResetPasswordSecret,
	// 	thisCodeFromRedis,
	// 	consts.UsersServiceName,
	// )
	// if !ok {
	// 	return thisLogicResponse
	// }
	//
	if paramRequest.Code != thisCodeFromRedis {
		thisMessage = append(thisMessage, messagepkg.Wrong(&thisLang, messagepkg.Keyword.Field["code"]))
		return responsepkg.Code401(
			thisMessage,
			paramRequest,
			responsepkg.EntityV1UseCaseFeature{
				Api: responsepkg.EntityV1UseCaseApi{
					WithMessage: true,
				},
			},
		)
	}

	// thisLogicResponse, thisToken, ok := logics.TokenCreate[string](
	thisLogicResponse, thisToken, ok := receiverUseCase.token.CreateString(
		paramRequest,
		thisLang,
		receiverUseCase.config.SecretKey.JwtResetPassword,
		ghafpkg.GenUUID(),
		receiverUseCase.config.ExpireDuration.In2m,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	// thisLogicResponse, ok = logics.RedisSet(
	thisLogicResponse, ok = receiverUseCase.redis.Set(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyAdminJWTResetPassword,
		thisMobileDetails.PhoneE164,
		thisToken,
		receiverUseCase.config.ExpireDuration.In2m,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	// SuccessFull
	thisMessage = append(thisMessage, messagepkg.Verify(&thisLang, messagepkg.Keyword.Section["admin_member"]))
	return responsepkg.Code200(
		thisMessage,
		response.UsersResetPasswordVerify{
			Mobile: thisMobileDetails.PhoneE164,
			Uuid:   thisToken,
		},
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
