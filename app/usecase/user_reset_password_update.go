package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersResetPasswordUpdate(
	paramCtx context.Context,
	paramRequest *request.UsersResetPasswordUpdate,
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

	thisLogicResponse, isExist, ok := receiverUseCase.redis.Exist(
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

	if !isExist {
		return responsepkg.Code401(
			thisMessage,
			nil, // Or paramRequest
			responsepkg.EntityV1UseCaseFeature{
				Api: responsepkg.EntityV1UseCaseApi{},
			},
		)
	}

	// thisLogicResponse, thisUUIDFromRedis, ok := logics.RedisGet(
	thisLogicResponse, thisUUIDFromRedis, ok := receiverUseCase.redis.Get(
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

	thisLogicResponse, thisTokenFromUUIDFromRedis, ok := receiverUseCase.token.ParseString(
		paramRequest,
		thisLang,
		receiverUseCase.config.SecretKey.JwtResetPassword,
		thisUUIDFromRedis,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	// thisLogicResponse, thisTokenFromUUIDRequest, ok := logics.TokenParse[string](
	thisLogicResponse, thisTokenFromUUIDRequest, ok := receiverUseCase.token.ParseString(
		paramRequest,
		thisLang,
		receiverUseCase.config.SecretKey.JwtResetPassword,
		paramRequest.Uuid,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	if thisTokenFromUUIDRequest.Data != thisTokenFromUUIDFromRedis.Data {
		thisMessage = append(thisMessage, messagepkg.NoAccess(&thisLang, messagepkg.Keyword.Field["no_access"]))
		return responsepkg.Code401(
			thisMessage,
			paramRequest,
			responsepkg.EntityV1UseCaseFeature{
				Api: responsepkg.EntityV1UseCaseApi{},
			},
		)
	}

	thisLogicResponse, ok = receiverUseCase.redis.Delete(
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

	thisCount, thisErr := receiverUseCase.model.UsersResetPasswordUpdate(
		paramCtx,
		thisMobileDetails.PhoneE164,
		paramRequest.NewPassword,
	)
	if thisErr != nil {
		thisErrorResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"reset_password_update",
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

	// SuccessFull
	thisMessage = append(thisMessage, messagepkg.Update(&thisLang, messagepkg.Keyword.Field["password"]))
	return responsepkg.Code200(
		thisMessage,
		nil, // Or response.UsersResetPasswordUpdate{}
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{},
		},
	)
}
