package myjwt

import (
	"fmt"
	"time"

	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/golang-jwt/jwt/v5"
)

type MyJWT struct {
}

func Init() *MyJWT {
	return &MyJWT{}
}

func (m MyJWT) CreateInt64(paramRequest interface{}, paramLang string, paramSecret string, paramValue int64, paramExpireTime time.Duration, paramOwnerName string) (responsepkg.EntityV1UseCase, string, bool) {
	return create[int64](
		paramRequest,
		paramLang,
		paramSecret,
		paramValue,
		paramExpireTime,
		paramOwnerName,
	)
}

func (m MyJWT) ParseInt64(paramRequest interface{}, paramLang string, paramSecret string, paramToken string, paramOwnerName string) (responsepkg.EntityV1UseCase, *_interface.CustomClaims[int64], bool) {
	return parse[int64](
		paramRequest,
		paramLang,
		paramSecret,
		paramToken,
		paramOwnerName,
	)
}

func (m MyJWT) CreateUint64(paramRequest interface{}, paramLang string, paramSecret string, paramValue uint64, paramExpireTime time.Duration, paramOwnerName string) (responsepkg.EntityV1UseCase, string, bool) {
	return create[uint64](
		paramRequest,
		paramLang,
		paramSecret,
		paramValue,
		paramExpireTime,
		paramOwnerName,
	)
}

func (m MyJWT) ParseUint64(paramRequest interface{}, paramLang string, paramSecret string, paramToken string, paramOwnerName string) (responsepkg.EntityV1UseCase, *_interface.CustomClaims[uint64], bool) {
	return parse[uint64](
		paramRequest,
		paramLang,
		paramSecret,
		paramToken,
		paramOwnerName,
	)
}

func (m MyJWT) ParseRawUint64(paramSecret string, paramToken string) (*_interface.CustomClaims[uint64], error) {
	return parseRawUint64[uint64](
		paramSecret,
		paramToken,
	)
}

func (m MyJWT) CreateString(paramRequest interface{}, paramLang string, paramSecret string, paramValue string, paramExpireTime time.Duration, paramOwnerName string) (responsepkg.EntityV1UseCase, string, bool) {
	return create[string](
		paramRequest,
		paramLang,
		paramSecret,
		paramValue,
		paramExpireTime,
		paramOwnerName,
	)
}

func (m MyJWT) ParseString(paramRequest interface{}, paramLang string, paramSecret string, paramToken string, paramOwnerName string) (responsepkg.EntityV1UseCase, *_interface.CustomClaims[string], bool) {
	return parse[string](
		paramRequest,
		paramLang,
		paramSecret,
		paramToken,
		paramOwnerName,
	)
}

func (m MyJWT) IsAuthorized(
	paramRequest interface{},
	paramLang string,
	paramSecret string,
	paramToken string,
	paramOwnerName string,
) (responsepkg.EntityV1UseCase, bool) {
	var thisMessage []messagepkg.Response

	_, err := jwt.Parse(paramToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(paramSecret), nil
	})

	if err != nil {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(
			&paramLang, "-logics_token_is_authorized"),
		)
		return responsepkg.Code500(
			thisMessage,
			responsepkg.Code500CustomResponseData{
				Request:          paramRequest,
				RequestOwnerName: paramOwnerName,
				Err:              err.Error(),
			},
		), false
	}

	return responsepkg.EntityV1UseCase{}, true
}

func (m MyJWT) IsAuthorizedRaw(
	paramSecret string,
	paramToken string,
) (bool, error) {
	_, err := jwt.Parse(paramToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(paramSecret), nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
