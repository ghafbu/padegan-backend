package _interface

import (
	"context"
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
)

type IUserSettingsModel interface {
	UserSettingsUpdate(
		paramCtx context.Context,
		paramRequest *request.UserSettingsUpdate,
	) (response.UserSettingsPublic, error)
	UserSettingsGet(paramCtx context.Context) (response.UserSettingsPublic, int64, error)
}
