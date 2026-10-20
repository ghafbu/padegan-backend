package myredis

import (
	"context"
	"fmt"
	"time"

	"github.com/ghafbu/padegan-backend/pkg/aespkg"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

func (r *MyRedis) SetWithEncrypt(
	paramCtx context.Context,
	paramLang string,
	paramRequest interface{},
	paramDefaultKey string,
	paramKey string,
	paramValue interface{},
	paramSecret string,
	paramExpireTime time.Duration,
	paramOwnerName string,
) (responsepkg.EntityV1UseCase, bool) {
	var thisMessage []messagepkg.Response

	thisCryptValue, thisErr := aespkg.New(paramSecret).Encrypt(fmt.Sprintf("%v", paramValue))
	if thisErr != nil {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "interfaceimplement_icache_myredis_set_with_encrypt_aes_encrypt"),
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

	thisErr = r.conn.Client().Set(
		paramCtx,
		ghafpkg.BuildKeyPattern(paramDefaultKey, paramKey),
		thisCryptValue,
		paramExpireTime,
	).Err()
	if thisErr != nil {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "interfaceimplement_icache_myredis_set_with_encrypt_redis_set"),
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
