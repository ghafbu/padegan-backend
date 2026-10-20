package model

import (
	"context"
	"errors"

	"github.com/ghafbu/padegan-backend/lib/consts"

	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/api/dto/response"
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/gormpkg"
)

func (myModel *UserModel) UsersReview(
	paramCtx context.Context,
	paramRequest *request.UsersReview,
) (response.UsersReview, int64, error) {
	// Variables
	var (
		thisResponse                    response.UsersReview
		thisUsers                       []response.UsersPublic
		thisOrderBy, thisOrderMethod    string
		thisPage, thisItemPerPage       int
		thisModifiedUsersPublicResponse = make([]response.UsersPublic, 0)
	)

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisResponse, 0, err
	}

	// DB Nil check
	if err := myModel.checkConnection(); err != nil {
		return thisResponse, 0, err
	}

	thisFetch := myModel.postgres.Conn().Model(&User{})

	if paramRequest.Name != nil && *paramRequest.Name != "" {
		thisFetch.Where("name=?", *paramRequest.Name)
	}

	if paramRequest.Family != nil && *paramRequest.Family != "" {
		thisFetch.Where("family=?", *paramRequest.Family)
	}

	if paramRequest.Mobile != nil && *paramRequest.Mobile != "" {
		thisFetch.Where("mobile=?", *paramRequest.Mobile)
	}

	if paramRequest.MobileCountryCode != nil && *paramRequest.MobileCountryCode != "" {
		thisFetch.Where("mobile_country_code=?", *paramRequest.MobileCountryCode)
	}

	if paramRequest.Gender != nil && *paramRequest.Gender != 0 {
		thisFetch.Where("gender=?", *paramRequest.Gender)
	}

	if paramRequest.Phone != nil && *paramRequest.Phone != "" {
		thisFetch.Where("phone=?", *paramRequest.Phone)
	}

	if paramRequest.PhoneCountryCode != nil && *paramRequest.PhoneCountryCode != "" {
		thisFetch.Where("phone_country_code=?", *paramRequest.PhoneCountryCode)
	}

	if paramRequest.BirthDate != nil && *paramRequest.BirthDate != "" {
		thisFetch.Where("birth_date=?", *paramRequest.BirthDate)
	}

	if paramRequest.NationalCode != nil && *paramRequest.NationalCode != "" {
		thisFetch.Where("national_code=?", *paramRequest.NationalCode)
	}

	if paramRequest.Email != nil && *paramRequest.Email != "" {
		thisFetch.Where("email=?", *paramRequest.Email)
	}

	if paramRequest.Username != nil && *paramRequest.Username != "" {
		thisFetch.Where("username=?", *paramRequest.Username)
	}

	if paramRequest.OrgThemeMode != nil && *paramRequest.OrgThemeMode != 0 {
		thisFetch.Where("org_theme_mode=?", *paramRequest.OrgThemeMode)
	}

	if paramRequest.Language != nil && *paramRequest.Language != 0 {
		thisFetch.Where("language=?", *paramRequest.Language)
	}

	if paramRequest.ProfileStatus != nil && *paramRequest.ProfileStatus != 0 {
		thisFetch.Where("profile_status=?", *paramRequest.ProfileStatus)
	}

	if paramRequest.Status != nil && *paramRequest.Status != 0 {
		thisFetch.Where("status=?", *paramRequest.Status)
	}

	// limit date [CreatedAt]
	if paramRequest.CreatedAtStart != nil && *paramRequest.CreatedAtStart != "" {
		if paramRequest.CreatedAtEnd != nil && *paramRequest.CreatedAtEnd != "" {
			thisFetch.Where("created_at BETWEEN ? AND ?", *paramRequest.CreatedAtStart, *paramRequest.CreatedAtEnd)
		} else {
			thisFetch.Where("created_at=?", *paramRequest.CreatedAtStart)
		}
	}

	// limit date [UpdatedAt]
	if paramRequest.UpdatedAtStart != nil && *paramRequest.UpdatedAtStart != "" {
		if paramRequest.UpdatedAtEnd != nil && *paramRequest.UpdatedAtEnd != "" {
			thisFetch.Where("updated_at BETWEEN ? AND ?", *paramRequest.UpdatedAtStart, *paramRequest.UpdatedAtEnd)
		} else {
			thisFetch.Where("updated_at=?", *paramRequest.UpdatedAtStart)
		}
	}

	// limit date [DeletedAt]
	// if paramRequest.DeletedAtStart != nil && *paramRequest.DeletedAtStart != "" {
	// 	if paramRequest.DeletedAtEnd != nil && *paramRequest.DeletedAtEnd != "" {
	// 		thisFetch.Where("deleted_at BETWEEN ? AND ?", *paramRequest.DeletedAtStart, *paramRequest.DeletedAtEnd)
	// 	} else {
	// 		thisFetch.Where("deleted_at=?", *paramRequest.DeletedAtStart)
	// 	}
	// }

	// pagination && sort
	thisPage = ghafpkg.GHIfNilOrDefaultTypeSetDefaultForINT(paramRequest.Page, 1)
	thisItemPerPage = ghafpkg.GHIfNilOrDefaultTypeSetDefaultForINT(paramRequest.ItemPerPage, consts.ItemPerPage10)
	thisOrderBy = ghafpkg.GHIfNilSetDefault[string](paramRequest.OrderBy, consts.DefaultModelOrderByID)
	thisOrderMethod = ghafpkg.GHIfNilSetDefault[string](paramRequest.OrderMethod, consts.DefaultModelOrderMethodASC)

	// Get TotalItems
	var thisTotalItems int64
	thisFetch.Count(&thisTotalItems)

	// DB Action
	thisFetch.
		Scopes(gormpkg.Paginate(thisPage, thisItemPerPage)).
		Order(thisOrderBy + " " + thisOrderMethod).
		Find(&thisUsers)

	if thisFetch.RowsAffected == 0 {
		return thisResponse, 0, errors.New(consts.ErrorKeyModelRecordNotFound)
	}

	if thisFetch.Error != nil {
		return thisResponse, thisFetch.RowsAffected, thisFetch.Error
	}

	for _, v := range thisUsers {
		if v.Birthdate != nil {
			ghafpkg.TimeGenerateStructDetailsWithPointer(v.Birthdate, &v.BirthdateDetails)
		}

		ghafpkg.TimeGenerateStructFullDetails(v.CreatedAt, &v.CreatedAtDetails)
		ghafpkg.TimeGenerateStructFullDetails(v.UpdatedAt, &v.UpdatedAtDetails)

		thisModifiedUsersPublicResponse = append(thisModifiedUsersPublicResponse, v)
	}

	// generate response
	thisResponse.Users = thisModifiedUsersPublicResponse
	thisResponse.Paper = response.UsersPaper{
		CurrentPage: thisPage,
		TotalPages:  ghafpkg.GetTotalPages(thisTotalItems, thisItemPerPage),
		TotalItems:  thisTotalItems,
		ItemPerPage: thisItemPerPage,
	}

	return thisResponse, thisFetch.RowsAffected, nil
}
