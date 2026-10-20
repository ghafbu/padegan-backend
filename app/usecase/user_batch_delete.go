package usecase

import (
	"context"

	"github.com/ghafbu/padegan-backend/api/dto/params"
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/lib/errors"
	"github.com/ghafbu/padegan-backend/lib/logics"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (receiverUseCase *UsersUseCase) UsersBatchDelete(
	paramCtx context.Context,
	paramRequest *request.UsersBatchDelete,
	paramHeader map[string][]string,
	paramParams params.UsersBatchDelete,
) responsepkg.EntityV1UseCase {
	thisLang := headerpkg.GetIndexAsString(headerpkg.ExistHeader(paramHeader, "S-Language"))
	var thisMessage []messagepkg.Response

	thisCount, thisErr := receiverUseCase.model.UsersBatchDelete(
		paramCtx,
		paramRequest,
		paramParams,
	)
	if thisErr != nil {
		thisLogicResponse := errors.ModelAction(
			thisErr,
			paramRequest,
			thisLang,
			"admin_members",
			"batch_delete",
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

	// Below code comment, because: 204 http response, not body returned for response
	// if thisCount == 1 {
	//	thisMessage = append(thisMessage, messagepkg.Delete(&thisLang, messagepkg.Keyword.Field["admin_member"]))
	// } else if thisCount > 1 {
	//	thisMessage = append(thisMessage, messagepkg.Delete(&thisLang, messagepkg.Keyword.Field["admin_members"]))
	// }

	// SuccessFull
	return responsepkg.Code204(
		thisMessage,
		nil, // Or any something
		responsepkg.EntityV1UseCaseFeature{
			Api: responsepkg.EntityV1UseCaseApi{},
		},
	)
}
