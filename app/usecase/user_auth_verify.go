package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersAuthVerify(
	paramCtx context.Context,
	paramRequest *request.UsersAuthVerify,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	thisLang := headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	// Logic: mobile details
	thisMobileDetails, thisLogicResponse, ok := logics.PhoneDetails(
		paramRequest, &paramRequest.Mobile, thisLang, "mobile",
	)
	if !ok {
		return thisLogicResponse
	}

	thisLogicResponse, thisCodeFromRedis, ok := receiverUseCase.redis.Get(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyOTP,
		thisMobileDetails.PhoneE164,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

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

	var requestForGet = request.UsersGet{Target: "mobile"}

	// Get user by mobile
	thisFetch, thisCount, thisErr := receiverUseCase.model.UsersGet(paramCtx, &requestForGet, thisMobileDetails.PhoneE164)
	if thisErr != nil {
		thisLogicResponse = errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"auth_verify",
		)
		return thisLogicResponse
	}

	// Logic: Not found
	thisLogicResponse, thisFound := logics.NotFound(
		paramRequest, thisCount, thisLang, "admin_member",
	)
	if !thisFound {
		return thisLogicResponse
	}

	// Create access token
	thisLogicResponse, thisAccessToken, ok := receiverUseCase.token.CreateUint64(
		paramRequest,
		thisLang,
		receiverUseCase.config.SecretKey.JwtAuthAccess,
		thisFetch.ID,
		receiverUseCase.config.ExpireDuration.In1w,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	thisLogicResponse, ok = receiverUseCase.redis.Set(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyAdminJWTAccess,
		// strconv.FormatUint(thisFetch.ID, 10), // thisMobileDetails.PhoneE164,
		thisMobileDetails.PhoneE164, // thisMobileDetails.PhoneE164,
		thisAccessToken,
		receiverUseCase.config.ExpireDuration.In1w,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	// Create refresh token
	// thisLogicResponse, thisRefreshToken, ok := logics.TokenCreate[uint64](
	thisLogicResponse, thisRefreshToken, ok := receiverUseCase.token.CreateUint64(
		paramRequest,
		thisLang,
		receiverUseCase.config.SecretKey.JwtAuthRefresh,
		thisFetch.ID,
		receiverUseCase.config.ExpireDuration.In1y,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	thisLogicResponse, ok = receiverUseCase.redis.Set(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyAdminJWTRefresh,
		thisMobileDetails.PhoneE164, // thisMobileDetails.PhoneE164,
		// strconv.FormatUint(thisFetch.ID, 10), // thisMobileDetails.PhoneE164,
		thisRefreshToken,
		receiverUseCase.config.ExpireDuration.In1y,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	// Successful
	thisMessage = append(thisMessage, messagepkg.Verify(&thisLang, messagepkg.Keyword.Section["admin_member"]))
	return responsepkg.Code200(
		thisMessage,
		response.UsersAuthVerify{
			User: thisFetch,
			Tokens: response.UsersToken{
				AccessToken:  thisAccessToken,
				RefreshToken: thisRefreshToken,
			},
		},
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{
				WithData: true,
			},
		},
	)
}
