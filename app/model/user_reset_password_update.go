package model

import (
	"context"
	"errors"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

func (myModel *UserModel) UsersResetPasswordUpdate(
	paramCtx context.Context,
	paramMobile string,
	paramNewPassword string,
) (int64, error) {
	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return 0, err
	}

	// hash password
	var HashPassword, thisErrPass = ghafpkg.BcryptPassword(paramNewPassword)
	if thisErrPass != nil {
		return 0, errors.New("password_hash_generate_error")
	}

	thisClone := User{
		Password:    &HashPassword,
		HasPassword: true,
	}

	thisFetch := myModel.postgres.Conn().Model(&thisClone).
		Where("mobile = ?", paramMobile).
		Updates(thisClone)

	if thisFetch.RowsAffected == 0 {
		return 0, errors.New("db_record_not_found")
	}

	if thisFetch.Error != nil {
		return thisFetch.RowsAffected, thisFetch.Error
	}

	return thisFetch.RowsAffected, nil
}
