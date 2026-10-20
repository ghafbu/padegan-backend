package myredis

import (
	"context"
	"errors"
	"fmt"
	"time"

	redispkg "github.com/ghafbu/padegan-backend/pkg/cachepkg/redis"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (r *MyRedis) SetAndGet(
	paramCtx context.Context,
	paramLang string,
	paramRequest interface{},
	paramDefaultKey string,
	paramKey string,
	paramValue interface{},
	paramExpireTime time.Duration,
	paramOwnerName string,
) (responsepkg.EntityV1UseCase, string, bool) {
	var (
		thisMessage     []messagepkg.Response
		thisKeyForRedis = ghafpkg.BuildKeyPattern(paramDefaultKey, paramKey)
	)

	thisErr := r.conn.Client().Set(
		paramCtx,
		thisKeyForRedis,
		paramValue,
		paramExpireTime,
	).Err()
	if thisErr != nil {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "interfaceimplement_icache_myredis_set_and_get_redis_set"),
		)
		return responsepkg.Code500( // internal error on redis
			thisMessage,
			responsepkg.Code500CustomResponseData{
				Request:          paramRequest,
				RequestOwnerName: paramOwnerName,
				Err:              thisErr.Error(),
			},
		), "", false
	}

	thisResult, thisErr := r.conn.Client().Get(paramCtx, thisKeyForRedis).Result()
	if thisErr != nil {
		if errors.Is(thisErr, redispkg.Nil) { // Not found on redis
			thisMessage = append(thisMessage, messagepkg.NotFoundSection(
				&paramLang,
				messagepkg.Keyword.Section[fmt.Sprintf("interfaceimplement_icache_myredis_set_and_get_redis_get_%s", thisKeyForRedis)]),
			)
			return responsepkg.Code404(
				thisMessage,
				paramRequest,
			), "", false
		}

		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "interfaceimplement_icache_myredis_set_and_get_redis_get"),
		)
		return responsepkg.Code500( // Internal error on redis
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
