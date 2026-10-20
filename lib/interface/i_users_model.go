package _interface

import (
	"context"

	"github.com/ghafbu/padegan-backend/api/dto/params"
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
)

type IUsersModel interface {
	UsersCreate(
		paramCtx context.Context,
		paramRequest *request.UsersCreate,
		paramMobileDetails phonepkg.PhoneDetails,
		paramPhoneDetails phonepkg.PhoneDetails,
	) (response.UsersPublic, error)
	UsersGet(
		paramCtx context.Context,
		paramRequest *request.UsersGet,
		paramId string,
	) (response.UsersPublic, int64, error)
	UsersUpdate(
		paramCtx context.Context,
		paramReq *request.UsersUpdate,
		paramPhoneDetails phonepkg.PhoneDetails,
		paramParams params.UsersUpdate,
	) (response.UsersPublic, int64, error)
	UsersReview(
		paramCtx context.Context,
		paramRequest *request.UsersReview,
	) (response.UsersReview, int64, error)
	UsersResetPasswordUpdate(
		paramCtx context.Context,
		paramNewPassword, paramMobile string,
	) (int64, error)
	UsersBatchGet(
		paramCtx context.Context,
		paramRequest *request.UsersBatchGet,
	) ([]response.UsersPublic, int64, error)
	UsersBatchUpdate(
		paramCtx context.Context,
		paramRequest *request.UsersBatchUpdate,
	) (response.UsersBatchUpdate, int64, error)
	UsersBatchDelete(
		paramCtx context.Context,
		paramRequest *request.UsersBatchDelete,
		paramParams params.UsersBatchDelete,
	) (int64, error)
}
