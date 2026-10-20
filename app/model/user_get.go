package model

import (
	"context"
	"errors"
	"strconv"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
)

func (myModel *UserModel) UsersGet(
	paramCtx context.Context,
	paramRequest *request.UsersGet,
	paramId string,
) (response.UsersPublic, int64, error) {
	//_, thisSpan := myModel.trace.Span(paramCtx, "DB Get")
	//thisSpan.SetAttributes(map[string]interface{}{
	//	"postgres.system":    "POSTGRES", // MYSQL, MARIA, POSTGRES, MONGO
	//	"postgres.operation": "SELECT",   // SELECT, INSERT, UPDATE, DELETE
	//})
	//defer thisSpan.End()

	var thisResponse response.UsersPublic

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisResponse, 0, err
	}

	thisFetch := myModel.postgres.Conn().Model(&User{})

	switch paramRequest.Target {
	case "mobile":
		thisFetch.Where("mobile=?", paramId)
	case "national_code":
		thisFetch.Where("national_code=?", paramId)
	case "email":
		thisFetch.Where("email=?", paramId)
	case "username":
		thisFetch.Where("username=?", paramId)
	default:
		thisId, err := strconv.Atoi(paramId)
		if err != nil || thisId == 0 {
			return thisResponse, 0, errors.New(consts.ErrorKeyModelIDBadRequest)
		}
		thisFetch.Where("id=?", thisId)
	}

	thisFetch.Find(&thisResponse)

	//thisSpan.SetAttribute("postgres.table", thisFetch.Statement.Table)

	if thisFetch.Error != nil {
		return thisResponse, thisFetch.RowsAffected, thisFetch.Error
	}

	if thisFetch.RowsAffected > 0 {
		if thisResponse.Birthdate != nil {
			ghafpkg.TimeGenerateStructDetailsWithPointer(thisResponse.Birthdate, &thisResponse.BirthdateDetails)
		}

		ghafpkg.TimeGenerateStructFullDetails(thisResponse.CreatedAt, &thisResponse.CreatedAtDetails)
		ghafpkg.TimeGenerateStructFullDetails(thisResponse.UpdatedAt, &thisResponse.UpdatedAtDetails)
	}

	return thisResponse, thisFetch.RowsAffected, nil
}
