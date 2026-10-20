package usecase

import (
	"context"
	"strconv"

	"github.com/ghafbu/padegan-backend/lib/logics"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersLogout(
	paramCtx context.Context,
	paramRequest *request.UsersLogout,
	paramHeader map[string][]string,
) responsepkg.EntityV1UseCase {
	thisLang := headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	var requestForGet = request.UsersGet{Target: "id"}
	thisFetch, thisCount, thisErr := receiverUseCase.model.UsersGet(
		paramCtx,
		&requestForGet,
		strconv.FormatUint(paramRequest.Id, 10))
	if thisErr != nil {
		thisErrorResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"logout",
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

	thisLogicResponse, ok := receiverUseCase.redis.Delete(
		paramCtx,
		thisLang,
		paramRequest,
		consts.DefaultKeyAdminJWTAccess,
		thisFetch.Mobile,
		receiverUseCase.config.Application.Name,
	)
	if !ok {
		return thisLogicResponse
	}

	// Successful
	thisMessage = append(thisMessage, messagepkg.Logout(&thisLang, messagepkg.Keyword.Section["admin_member"]))
	return responsepkg.Code204(
		thisMessage,
		nil,
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{},
		},
	)
}
