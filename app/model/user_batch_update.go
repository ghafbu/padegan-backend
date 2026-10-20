package model

import (
	"context"
	"errors"
	"time"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/gormpkg"
	"gorm.io/gorm"
)

func (myModel *UserModel) UsersBatchUpdate(
	paramCtx context.Context,
	paramRequest *request.UsersBatchUpdate,
) (response.UsersBatchUpdate, int64, error) {
	var (
		thisResponse           response.UsersBatchUpdate
		thisPublicResponse     []response.UsersPublic
		thisDefaultPage        = consts.DefaultPage
		thisDefaultItemPerPage = consts.ItemPerPage10
		thisDefaultOrderBy     = consts.DefaultModelOrderByID
		thisDefaultOrderMethod = consts.DefaultModelOrderMethodASC
	)

	thisClone := User{
		Status:    StatusUser(paramRequest.Status),
		UpdatedAt: time.Now(),
	}

	// fetch
	thisFetch := myModel.postgres.Conn().Session(&gorm.Session{SkipHooks: true}).Model(&User{})

	// update
	thisFetch.
		Where("id IN ?", paramRequest.Ids).
		Updates(thisClone)

	// err-update
	if thisFetch.Error != nil {
		return thisResponse, thisFetch.RowsAffected, thisFetch.Error
	}

	// get TotalItems
	var thisTotalItems int64
	thisFetch.Count(&thisTotalItems)

	if paramRequest.Page != nil && *paramRequest.Page != 0 {
		thisDefaultPage = *paramRequest.Page
	}

	// get ItemPerPage5 EX: 20 | 50 | 100 ....
	if paramRequest.ItemPerPage != nil && *paramRequest.ItemPerPage != 0 {
		thisDefaultItemPerPage = *paramRequest.ItemPerPage
	}
	var thisItemPerPage int = ghafpkg.GetItemPerPage(thisDefaultItemPerPage, consts.ItemPerPage5)

	if paramRequest.OrderBy != nil && *paramRequest.OrderBy != "" {
		thisDefaultOrderBy = *paramRequest.OrderBy
	}

	if paramRequest.OrderMethod != nil && *paramRequest.OrderMethod != "" {
		thisDefaultOrderMethod = *paramRequest.OrderMethod
	}

	// find
	thisFetch.
		Scopes(gormpkg.Paginate(thisDefaultPage, thisItemPerPage)).
		Order(thisDefaultOrderBy + " " + thisDefaultOrderMethod).
		Find(&thisPublicResponse)

	if thisFetch.RowsAffected == 0 {
		return thisResponse, 0, errors.New(consts.ErrorKeyModelRecordNotFound)
	}

	if thisFetch.Error != nil {
		return thisResponse, thisFetch.RowsAffected, thisFetch.Error
	}

	var thisModifiedPublicResponse = make([]response.UsersPublic, 0)
	for _, v := range thisPublicResponse {
		if v.Birthdate != nil {
			ghafpkg.TimeGenerateStructDetailsWithPointer(v.Birthdate, &v.BirthdateDetails)
		}
		ghafpkg.TimeGenerateStructFullDetails(v.CreatedAt, &v.CreatedAtDetails)
		ghafpkg.TimeGenerateStructFullDetails(v.UpdatedAt, &v.UpdatedAtDetails)

		thisModifiedPublicResponse = append(thisModifiedPublicResponse, v)
	}

	thisResponse.Users = thisModifiedPublicResponse
	thisResponse.Paper = response.UsersPaper{
		CurrentPage: thisDefaultPage,
		TotalPages:  ghafpkg.GetTotalPages(thisTotalItems, thisItemPerPage),
		TotalItems:  thisTotalItems,
		ItemPerPage: thisItemPerPage,
	}

	return thisResponse, thisFetch.RowsAffected, nil
}
