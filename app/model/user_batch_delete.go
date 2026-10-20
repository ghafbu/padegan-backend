package model

import (
	"context"

	"github.com/ghafbu/padegan-backend/api/dto/params"
	"github.com/ghafbu/padegan-backend/api/dto/request"
)

func (myModel *UserModel) UsersBatchDelete(
	paramCtx context.Context,
	paramRequest *request.UsersBatchDelete,
	paramParams params.UsersBatchDelete,
) (int64, error) {
	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return 0, err
	}

	// postgres.Delete(&users, []int{1,2,3}) -> DELETE FROM users WHERE id IN (1,2,3)

	thisClone := User{}

	thisFetch := myModel.postgres.Conn().Delete(&thisClone, paramRequest.Ids)

	if thisFetch.Error != nil {
		return 0, thisFetch.Error
	}

	return thisFetch.RowsAffected, nil
}
