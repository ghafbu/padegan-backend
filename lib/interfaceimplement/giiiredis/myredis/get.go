package myredis

import (
	"context"
	"errors"
	"fmt"

	redispkg "github.com/ghafbu/padegan-backend/pkg/cachepkg/redis"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (r *MyRedis) Get(
	paramCtx context.Context,
	paramLang string,
	paramRequest interface{},
	paramDefaultKey string,
	paramKey string,
	paramOwnerName string,
) (responsepkg.EntityV1UseCase, string, bool) {
	var (
		thisMessage     []messagepkg.Response
		thisKeyForRedis = ghafpkg.BuildKeyPattern(paramDefaultKey, paramKey)
	)

	thisResult, thisErr := r.conn.Client().Get(paramCtx, thisKeyForRedis).Result()
	if thisErr != nil {
		if errors.Is(thisErr, redispkg.Nil) { // Not found
			thisMessage = append(thisMessage, messagepkg.NotFoundSection(
				&paramLang,
				messagepkg.Keyword.Section[fmt.Sprintf("interfaceimplement_icache_myredis_get_%s", thisKeyForRedis)]),
			)
			return responsepkg.Code404(
				thisMessage,
				paramRequest,
			), "", false
		}

		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "interfaceimplement_icache_myredis_get"),
		)
		return responsepkg.Code500(
			thisMessage,
			responsepkg.Code500CustomResponseData{
				Request:          paramRequest,
				RequestOwnerName: paramOwnerName,
				Err:              thisErr.Error(),
			},
		), "", false
	}

	return responsepkg.EntityV1UseCase{}, thisResult, true
}
