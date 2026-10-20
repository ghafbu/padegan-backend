package myredis

import (
	"context"
	"time"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (r *MyRedis) Set(
	paramCtx context.Context,
	paramLang string,
	paramRequest interface{},
	paramDefaultKey string,
	paramKey string,
	paramValue interface{},
	paramExpireTime time.Duration,
	paramOwnerName string,
) (responsepkg.EntityV1UseCase, bool) {
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
			&paramLang, "interfaceimplement_icache_myredis_set"),
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

	return responsepkg.EntityV1UseCase{}, true
}
