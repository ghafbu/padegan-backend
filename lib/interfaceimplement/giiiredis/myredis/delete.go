package myredis

import (
	"context"
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (r *MyRedis) Delete(
	paramCtx context.Context,
	paramLang string,
	paramRequest interface{},
	paramDefaultKey string,
	paramKey string,
	paramOwnerName string,
) (responsepkg.EntityV1UseCase, bool) {
	var (
		thisMessage     []messagepkg.Response
		thisKeyForRedis = ghafpkg.BuildKeyPattern(paramDefaultKey, paramKey)
	)

	thisResult, thisErr := r.conn.Client().Del(paramCtx, thisKeyForRedis).Result()
	if thisErr != nil { // Internal error on redis
		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "interfaceimplement_icache_myredis_delete"),
		)
		return responsepkg.Code500(
			thisMessage,
			responsepkg.Code500CustomResponseData{
				Request:          paramRequest,
				RequestOwnerName: paramOwnerName,
				Err:              thisErr.Error(),
			},
		), false
	}

	if !(thisResult != 0) { // Not found key on redis
		thisMessage = append(thisMessage, messagepkg.NotFoundSection(
			&paramLang,
			messagepkg.Keyword.Section[fmt.Sprintf("interfaceimplement_icache_myredis_delete_%s", thisKeyForRedis)]),
		)
		return responsepkg.Code404(
			thisMessage,
			paramRequest,
		), false
	}

	return responsepkg.EntityV1UseCase{}, true
}
