package _interface

import (
	"time"

	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId uint64 `json:"id"`
}

type CustomClaims[T any] struct {
	jwt.RegisteredClaims
	Data T
}

type IToken interface {
	CreateInt64(
		paramRequest interface{},
		paramLang string,
		paramSecret string,
		paramValue int64,
		paramExpireTime time.Duration,
		paramOwnerName string,
	) (responsepkg.EntityV1UseCase,
		string,
		bool,
	)

	ParseInt64(
		paramRequest interface{},
		paramLang string,
		paramSecret string,
		paramToken string,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		*CustomClaims[int64],
		bool,
	)

	CreateUint64(
		paramRequest interface{},
		paramLang string,
		paramSecret string,
		paramValue uint64,
		paramExpireTime time.Duration,
		paramOwnerName string,
	) (responsepkg.EntityV1UseCase,
		string,
		bool,
	)

	ParseUint64(
		paramRequest interface{},
		paramLang string,
		paramSecret string,
		paramToken string,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		*CustomClaims[uint64],
		bool,
	)

	ParseRawUint64(
		paramSecret string,
		paramToken string,
	) (
		*CustomClaims[uint64],
		error,
	)

	CreateString(
		paramRequest interface{},
		paramLang string,
		paramSecret string,
		paramValue string,
		paramExpireTime time.Duration,
		paramOwnerName string,
	) (responsepkg.EntityV1UseCase,
		string,
		bool)

	ParseString(
		paramRequest interface{},
		paramLang string,
		paramSecret string,
		paramToken string,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		*CustomClaims[string],
		bool,
	)

	IsAuthorized(
		paramRequest interface{},
		paramLang string,
		paramSecret string,
		paramToken string,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		bool,
	)
	IsAuthorizedRaw(
		paramSecret string,
		paramToken string,
	) (
		bool,
		error,
	)
}
