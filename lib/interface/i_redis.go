package _interface

import (
	"context"
	"time"

	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
)

type IRedis interface {
	Get(
		paramCtx context.Context,
		paramLang string,
		paramRequest interface{},
		paramDefaultKey string,
		paramKey string,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		string,
		bool,
	)
	ExistAndGetRaw(
		paramCtx context.Context,
		paramDefaultKey string,
		paramKey string,
	) (
		string,
		bool,
		error,
	)
	GetWithDecrypt(
		paramCtx context.Context,
		paramLang string,
		paramRequest interface{},
		paramDefaultKey string,
		paramKey string,
		paramSecret string,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		string,
		bool,
	)
	Set(
		paramCtx context.Context,
		paramLang string,
		paramRequest interface{},
		paramDefaultKey string,
		paramKey string,
		paramValue interface{},
		paramExpireTime time.Duration,
		paramOwnerName string,
	) (responsepkg.EntityV1UseCase, bool)
	SetWithEncrypt(
		paramCtx context.Context,
		paramLang string,
		paramRequest interface{},
		paramDefaultKey string,
		paramKey string,
		paramValue interface{},
		paramSecret string,
		paramExpireTime time.Duration,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		bool,
	)
	SetAndGet(
		paramCtx context.Context,
		paramLang string,
		paramRequest interface{},
		paramDefaultKey string,
		paramKey string,
		paramValue interface{},
		paramExpireTime time.Duration,
		paramOwnerName string,
	) (
		responsepkg.EntityV1UseCase,
		string,
		bool,
	)
	Exist(
		paramCtx context.Context,
		paramLang string,
		paramRequest interface{},
		paramDefaultKey string,
		paramKey string,
		paramOwnerName string,
	) (responsepkg.EntityV1UseCase, bool, bool)
	Delete(
		paramCtx context.Context,
		paramLang string,
		paramRequest interface{},
		paramDefaultKey string,
		paramKey string,
		paramOwnerName string,
	) (responsepkg.EntityV1UseCase, bool)
	Close() error
}
